package config

import (
	"github.com/cloudwego/kitex/pkg/connpool"
	"time"
)

var LongConnConfig = connpool.IdleConfig{
	MaxIdlePerAddress: 10,
	MaxIdleGlobal:     100,
	MaxIdleTimeout:    time.Minute,
	MinIdlePerAddress: 2,
}
