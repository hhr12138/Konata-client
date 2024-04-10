package Konata_client

import (
	"fmt"
	"time"
)

type Cmder interface {
	Args() []interface{}
	arg(int) string
	Name() string

	//readReply() error
	SetErr(error)

	readTimeout() *time.Duration

	Err() error
	fmt.Stringer
}
