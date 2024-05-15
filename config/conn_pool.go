package config

import (
	"github.com/cloudwego/kitex/pkg/connpool"
	"time"
)

var LongConnConfig = connpool.IdleConfig{
	MaxIdlePerAddress: 20,
	MaxIdleGlobal:     100,
	MaxIdleTimeout:    time.Minute,
	MinIdlePerAddress: 3,
}
