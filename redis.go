package Konata_client

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/hhr12138/Konata-client/config"
	"github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client"
	"github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client/konataservice"
	"github.com/hhr12138/Konata-client/utils"
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
	Addr atomic.Value
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
	c.Addr.Store(config.DefaultAddrs[0])
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
		resp  *konata_client.Reply
		opts  []callopt.Option
	)
	ctx = context.WithValue(ctx, "req_id", reqId)
	for attempt := 0; ; attempt++ {
		opts = c.buildOpt()
		if attempt > 0 {
			// 睡眠一段时间，等待网络恢复
			time.Sleep(c.retryBackoff(attempt))
		}

		command, err := utils.BuildCommand(cmd.Args())
		if err != nil {
			cmd.SetErr(err)
			return err
		}
		if cmd.GetOp() == konata_client.Write {
			putAppendArgs := &konata_client.PutAppendArgs_{
				ReqId:   reqId,
				Command: command,
				Op:      cmd.GetOp(),
			}
			resp, err = c.kitexClient.PutAppend(ctx, putAppendArgs, opts...)
		} else {
			getArgs := &konata_client.GetArgs_{
				ReqId:   reqId,
				Command: command,
				Op:      cmd.GetOp(),
			}
			resp, err = c.kitexClient.Get(ctx, getArgs)
		}
		// 网络异常/重试异常
		if err != nil || (resp.Error != nil && resp.Error.Repeat) {
			// 当前服务不是master，更换master.
			if resp.Error.Code == konata_client.ErrCodeMasterReplace {
				c.Addr.Store(resp.Base.Addr)
			}
			continue
		}
		// 明确执行失败的异常直接抛出
		if resp.Error != nil {
			err = fmt.Errorf("err_code=%v,message=%v", resp.Error.Code, resp.Error.Message)
			cmd.SetErr(err)
			return err
		}
	}
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

func (c *DefaultClient) buildOpt() []callopt.Option {
	opts := make([]callopt.Option, 0)
	val := c.Addr.Load()
	// 不判ok，这里绝对不能错。
	addr := val.(string)
	opts = append(opts, callopt.WithHTTPHost(addr))
	return opts
}
