package Konata_client

import (
	"github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client"
	"time"
)

type Cmder interface {
	Args() []interface{}
	arg(int) string
	Name() string

	//readReply() error
	SetErr(error)

	SetOp(op konata_client.OpType)
	GetOp() konata_client.OpType

	readReply(command string) error

	readTimeout() *time.Duration

	Err() error
}
