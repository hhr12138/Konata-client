package Konata_client

import (
	"github.com/hhr12138/Konata-client/internal"
	"github.com/hhr12138/Konata-client/internal/proto"
	"github.com/hhr12138/Konata-client/internal/util"
	"github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client"
	"strconv"
	"strings"
	"time"
)

type Cmdable interface {
	StringCmdable
	ListCmdable
	HashCmdable
	SetCmdable
}

type process func(cmd Cmder) error

type baseCmd struct {
	_args []interface{}
	Op    konata_client.OpType
	err   error

	_readTimeout *time.Duration
}

func (cmd *baseCmd) ReadReply(command string) error {
	return nil
}

func (cmd *baseCmd) Err() error {
	return cmd.err
}

func (cmd *baseCmd) Args() []interface{} {
	return cmd._args
}

func (cmd *baseCmd) arg(pos int) string {
	if pos < 0 || pos >= len(cmd._args) {
		return ""
	}
	s, _ := cmd._args[pos].(string)
	return s
}

func (cmd *baseCmd) SetOp(op konata_client.OpType) {
	cmd.Op = op
}

func (cmd *baseCmd) GetOp() konata_client.OpType {
	return cmd.Op
}

func (cmd *baseCmd) Name() string {
	if len(cmd._args) > 0 {
		// Cmd name must be lower cased.
		s := strings.ToLower(cmd.arg(0))
		cmd._args[0] = s
		return s
	}
	return ""
}

func (cmd *baseCmd) readTimeout() *time.Duration {
	return cmd._readTimeout
}

func (cmd *baseCmd) setReadTimeout(d time.Duration) {
	cmd._readTimeout = &d
}

func (cmd *baseCmd) ReadTimeout() *time.Duration {
	return cmd.readTimeout()
}

func (cmd *baseCmd) SetReadTimeout(d time.Duration) {
	cmd.setReadTimeout(d)
}

func (cmd *baseCmd) SetErr(e error) {
	cmd.err = e
}

//------------------------------------------------------------------------------

type IntCmd struct {
	baseCmd

	val int64
}

func NewIntCmd(args ...interface{}) *IntCmd {
	return &IntCmd{
		baseCmd: baseCmd{_args: args},
	}
}

func (cmd *IntCmd) Val() int64 {
	return cmd.val
}

func (cmd *IntCmd) SetVal(val int64) {
	cmd.val = val
}

func (cmd *IntCmd) Result() (int64, error) {
	return cmd.val, cmd.err
}

func (cmd *IntCmd) String() string {
	return cmdString(cmd, cmd.val)
}

//------------------------------------------------------------------------------

type StringCmd struct {
	baseCmd

	val []byte
}

func NewStringCmd(args ...interface{}) *StringCmd {
	return &StringCmd{
		baseCmd: baseCmd{_args: args},
	}
}

func (cmd *StringCmd) Val() string {
	return string(cmd.val)
}

func (cmd *StringCmd) SetVal(v string) {
	cmd.val = []byte(v)
}

func (cmd *StringCmd) Result() (string, error) {
	return cmd.Val(), cmd.err
}

func (cmd *StringCmd) Bytes() ([]byte, error) {
	return cmd.val, cmd.err
}

func (cmd *StringCmd) Int64() (int64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseInt(cmd.Val(), 10, 64)
}

func (cmd *StringCmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseUint(cmd.Val(), 10, 64)
}

func (cmd *StringCmd) Float64() (float64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseFloat(cmd.Val(), 64)
}

func (cmd *StringCmd) Scan(val interface{}) error {
	if cmd.err != nil {
		return cmd.err
	}
	return proto.Scan(cmd.val, val)
}

func (cmd *StringCmd) String() string {
	return cmdString(cmd, cmd.val)
}

func cmdString(cmd Cmder, val interface{}) string {
	b := make([]byte, 0, 64)

	for i, arg := range cmd.Args() {
		if i > 0 {
			b = append(b, ' ')
		}
		b = internal.AppendArg(b, arg)
	}

	if err := cmd.Err(); err != nil {
		b = append(b, ": "...)
		b = append(b, err.Error()...)
	} else if val != nil {
		b = append(b, ": "...)
		b = internal.AppendArg(b, val)
	}

	return util.BytesToString(b)
}

//------------------------------------------------------------------------------

type FloatCmd struct {
	baseCmd

	val float64
}

func NewFloatCmd(args ...interface{}) *FloatCmd {
	return &FloatCmd{
		baseCmd: baseCmd{_args: args},
	}
}

func (cmd *FloatCmd) Val() float64 {
	return cmd.val
}

func (cmd *FloatCmd) SetVal(val float64) {
	cmd.val = val
}

func (cmd *FloatCmd) Result() (float64, error) {
	return cmd.Val(), cmd.Err()
}

func (cmd *FloatCmd) String() string {
	return cmdString(cmd, cmd.val)
}

//------------------------------------------------------------------------------

type SliceCmd struct {
	baseCmd

	val []interface{}
}

func NewSliceCmd(args ...interface{}) *SliceCmd {
	return &SliceCmd{
		baseCmd: baseCmd{_args: args},
	}
}

func (cmd *SliceCmd) Val() []interface{} {
	return cmd.val
}

func (cmd *SliceCmd) SetVal(val []interface{}) {
	cmd.val = val
}

func (cmd *SliceCmd) Result() ([]interface{}, error) {
	return cmd.val, cmd.err
}

func (cmd *SliceCmd) String() string {
	return cmdString(cmd, cmd.val)
}

//------------------------------------------------------------------------------

type StatusCmd struct {
	baseCmd

	val string
}

func NewStatusCmd(args ...interface{}) *StatusCmd {
	return &StatusCmd{
		baseCmd: baseCmd{_args: args},
	}
}

func (cmd *StatusCmd) Val() string {
	return cmd.val
}

func (cmd *StatusCmd) SetVal(val string) {
	cmd.val = val
}

func (cmd *StatusCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

func (cmd *StatusCmd) String() string {
	return cmdString(cmd, cmd.val)
}

//------------------------------------------------------------------------------

type BoolCmd struct {
	baseCmd

	val bool
}

func NewBoolCmd(args ...interface{}) *BoolCmd {
	return &BoolCmd{
		baseCmd: baseCmd{_args: args},
	}
}

func (cmd *BoolCmd) Val() bool {
	return cmd.val
}

func (cmd *BoolCmd) SetVal(val bool) {
	cmd.val = val
}

func (cmd *BoolCmd) Result() (bool, error) {
	return cmd.val, cmd.err
}

func (cmd *BoolCmd) String() string {
	return cmdString(cmd, cmd.val)
}

var ok = []byte("OK")

//------------------------------------------------------------------------------

type StringSliceCmd struct {
	baseCmd

	val []string
}

func NewStringSliceCmd(args ...interface{}) *StringSliceCmd {
	return &StringSliceCmd{
		baseCmd: baseCmd{_args: args},
	}
}

func (cmd *StringSliceCmd) Val() []string {
	return cmd.val
}

func (cmd *StringSliceCmd) SetVal(val []string) {
	cmd.val = val
}

func (cmd *StringSliceCmd) Result() ([]string, error) {
	return cmd.Val(), cmd.Err()
}

func (cmd *StringSliceCmd) String() string {
	return cmdString(cmd, cmd.val)
}

func (cmd *StringSliceCmd) ScanSlice(container interface{}) error {
	return proto.ScanSlice(cmd.Val(), container)
}

//------------------------------------------------------------------------------

type StringStringMapCmd struct {
	baseCmd

	val map[string]string
}

func NewStringStringMapCmd(args ...interface{}) *StringStringMapCmd {
	return &StringStringMapCmd{
		baseCmd: baseCmd{_args: args},
	}
}

func (cmd *StringStringMapCmd) Val() map[string]string {
	return cmd.val
}

func (cmd *StringStringMapCmd) SetVal(val map[string]string) {
	cmd.val = val
}

func (cmd *StringStringMapCmd) Result() (map[string]string, error) {
	return cmd.val, cmd.err
}

func (cmd *StringStringMapCmd) String() string {
	return cmdString(cmd, cmd.val)
}
