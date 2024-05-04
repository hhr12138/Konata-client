package Konata_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/hhr12138/Konata-client/config"
	"github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client"
	"github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client/konataservice"
	"github.com/hhr12138/Konata-client/utils"
	"github.com/pkg/errors"
	"math/rand"
	"sync/atomic"
	"time"
)

type DefaultClient struct {
	// 只读
	opt *Options
	// 只读
	process
	// 只读
	kitexClient konataservice.Client
	// 注意并发
	Addrs []atomic.Value
}

// NewClient returns a client to the Redis Server specified by Options.
func NewClient(opt *Options) (*DefaultClient, error) {
	opt.init()

	c := DefaultClient{
		opt: opt,
	}
	err := c.init()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *DefaultClient) init() error {
	kitexClient, err := konataservice.NewClient(config.PSM, c.buildKitexOpt()...)
	if err != nil {
		return err
	}
	c.kitexClient = kitexClient
	c.process = c.defaultProcess
	c.Addrs = make([]atomic.Value, len(config.DefaultAddrs))
	for i, addr := range config.DefaultAddrs {
		c.Addrs[i].Store(addr)
	}
	return nil
}

func (c *DefaultClient) buildKitexOpt() []client.Option {
	opts := make([]client.Option, 0)
	// build addrs
	opts = append(opts, client.WithHostPorts(config.DefaultAddrs...))
	opts = append(opts, client.WithLongConnection(config.LongConnConfig))
	opts = append(opts, client.WithConnectTimeout(c.opt.DialTimeout))
	opts = append(opts, client.WithRPCTimeout(c.opt.RPCTimeout))
	return opts
}

func (c *DefaultClient) defaultProcess(cmd Cmder) error {
	var (
		ctx = context.Background()
		// 生成唯一id
		reqId = utils.GetReqId()
		// 必须为指针，从而让defer函数能感知到变化
		targetAddr *string
		resp       *konata_client.Reply
		opts       []callopt.Option
		key        = cmd.Args()[1].(string)
		// 根据keyhash 得到处理他的节点
		addrIdx = utils.GetAddrIdx(key)
	)
	ctx = context.WithValue(ctx, "req_id", reqId)
	defer c.RemoveReqId(ctx, reqId, targetAddr)
	for attempt := 0; ; attempt++ {
		val := c.Addrs[addrIdx].Load()
		// 不判ok，这里绝对不能错。
		addr := val.(string)
		targetAddr = &addr
		opts = c.buildOpt(addr)
		if attempt > 0 {
			// 睡眠一段时间，等待网络恢复
			time.Sleep(c.retryBackoff(attempt))
		}

		msg, err := utils.BuildMsg(cmd.Args())
		if err != nil {
			cmd.SetErr(err)
			return err
		}
		command := c.buildCommand(reqId, msg)
		commandBs, _ := json.Marshal(command)
		if cmd.GetOp() == konata_client.Write {
			putAppendArgs := &konata_client.PutAppendArgs_{
				ReqId:   reqId,
				Command: string(commandBs),
				Op:      cmd.GetOp(),
			}
			resp, err = c.kitexClient.PutAppend(ctx, putAppendArgs, opts...)
		} else {
			getArgs := &konata_client.GetArgs_{
				ReqId:   reqId,
				Command: string(commandBs),
				Op:      cmd.GetOp(),
			}
			resp, err = c.kitexClient.Get(ctx, getArgs, opts...)
		}
		// 网络异常/重试异常
		if err != nil || (resp.Error != nil && resp.Error.Repeat) {
			// 当前服务不是master，更换master.
			if resp.Error.Code == konata_client.ErrCodeMasterReplace {
				c.Addrs[addrIdx].Store(resp.Base.Addr)
			}
			continue
		}
		// 明确执行失败的异常直接抛出
		if resp.Error != nil {
			err = fmt.Errorf("err_code=%v,message=%v", resp.Error.Code, resp.Error.Message)
			cmd.SetErr(err)
			return err
		}
		err = cmd.readReply(resp.Value)
		if err != nil {
			cmd.SetErr(errors.Wrapf(err, "err_code=%v,返回值解析失败", konata_client.ErrCodeRspParseFail))
			return err
		}
		return nil
	}
}

// 执行结束后异步删除req_id防止oom，服务端后等待一个MSL后删除
func (c *DefaultClient) RemoveReqId(ctx context.Context, reqId string, targetAddr *string) {
	// 理论上不会为空，防御性编程
	if targetAddr == nil {
		return
	}
	// 一直重试直到成功，防止oom
	go func() {
		opts := c.buildOpt(*targetAddr)
		for {
			rsp, err := c.kitexClient.RemoveReqId(context.Background(), &konata_client.GetArgs_{ReqId: reqId}, opts...)
			// 只要不是语句解析错误，那么就该重试
			if err != nil || (rsp.Error != nil && rsp.Error.Code != konata_client.ErrCodeCommandParseFail) {
				continue
			}
			break
		}
	}()
}

func (c *DefaultClient) retryBackoff(retry int) time.Duration {
	if retry < 0 {
		retry = 0
	}

	backoff := c.opt.MinRetryBackoff << uint(retry)
	// 处理溢出情况
	if backoff > c.opt.MaxRetryBackoff || backoff < c.opt.MinRetryBackoff {
		backoff = c.opt.MaxRetryBackoff
	}

	if backoff == 0 {
		return 0
	}
	// 随机重试时间
	return time.Duration(rand.Int63n(int64(backoff)))
}

func (c *DefaultClient) buildOpt(addr string) []callopt.Option {
	opts := make([]callopt.Option, 0)
	opts = append(opts, callopt.WithHTTPHost(addr))
	return opts
}

func (c *DefaultClient) buildCommand(reqId, msg string) *konata_client.Command {
	return &konata_client.Command{
		ReqId: reqId,
		Msg:   msg,
	}
}
