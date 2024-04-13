package Konata_client

import (
	"time"
)

// 为了确保日后rpc框架切换对用户无感知，必须包一层自己对options
type Options struct {
	// host:port address.
	// Konata在客户端视角为单机服务，想用户屏蔽具体addr
	//Addr string

	// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
	ClientName string

	//无限期重试，直到成功
	//MaxRetries int
	MaxRetryBackoff time.Duration
	MinRetryBackoff time.Duration

	DialTimeout time.Duration
	RPCTimeout  time.Duration

	readOnly bool
}

func (opt *Options) init() {
	//if opt.Addr == "" {
	//	opt.Addr = config.DefaultAddr
	//}
	if opt.DialTimeout == 0 {
		opt.DialTimeout = 5 * time.Second
	}
	switch opt.RPCTimeout {
	case -1:
		opt.RPCTimeout = 0
	case 0:
		opt.RPCTimeout = 3 * time.Second
	}
	switch opt.MinRetryBackoff {
	case -1:
		opt.MinRetryBackoff = 0
	case 0:
		opt.MinRetryBackoff = 8 * time.Millisecond
	}
	switch opt.MaxRetryBackoff {
	case -1:
		opt.MaxRetryBackoff = 0
	case 0:
		opt.MaxRetryBackoff = 512 * time.Millisecond
	}
}
