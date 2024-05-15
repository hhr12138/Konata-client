package Konata_client

import "github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client"

type StringCmdable interface {
	// string
	Append(key, value string) *IntCmd
	Decr(key string) *IntCmd
	DecrBy(key string, decrement int64) *IntCmd
	DecrByFloat(key string, decrement float64)
	Get(key string) *StringCmd
	GetDel(keys ...string) *StringCmd
	GetRange(key string, start, end int64) *StringCmd
	GetSet(key string, value interface{}) *StringCmd
	Incr(key string) *IntCmd
	IncrBy(key string, value int64) *IntCmd
	IncrByFloat(key, incr float64) *FloatCmd
	MGet(keys ...string) *SliceCmd
	MSet(pairs ...interface{}) *StatusCmd
	Set(key string, value interface{}) *StatusCmd
	StrLen(key string) *IntCmd
}

func (p process) Append(key, value string) *IntCmd {
	cmd := NewIntCmd("append", key, value)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func (p process) Decr(key string) *IntCmd {
	cmd := NewIntCmd("decr", key)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func (p process) DecrBy(key string, decrement int64) *IntCmd {
	cmd := NewIntCmd("decrby", key, decrement)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}
func (p process) DecrByFloat(key string, value float64) *FloatCmd {
	cmd := NewFloatCmd("decrbyfloat", key, value)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func (p process) GetRange(key string, start, end int64) *StringCmd {
	cmd := NewStringCmd("getrange", key, start, end)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

// GetDel redis-server version >= 6.2.0.
func (p process) GetDel(key string) *StringCmd {
	cmd := NewStringCmd("getdel", key)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

// Get Redis `GET key` command. It returns redis.Nil error when key does not exist.
func (p process) Get(key string) *StringCmd {
	cmd := NewStringCmd("get", key)
	cmd.SetOp(konata_client.Read)
	_ = p(cmd)
	return cmd
}

func (p process) GetSet(key string, value interface{}) *StringCmd {
	cmd := NewStringCmd("getset", key, value)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func (p process) Incr(key string) *IntCmd {
	cmd := NewIntCmd("incr", key)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func (p process) IncrBy(key string, value int64) *IntCmd {
	cmd := NewIntCmd("incrby", key, value)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func (p process) IncrByFloat(key string, value float64) *FloatCmd {
	cmd := NewFloatCmd("incrbyfloat", key, value)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func (p process) MGet(keys ...string) *SliceCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = "mget"
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewSliceCmd(args...)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

// MSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
//   - MSet(struct), For struct types, see HSet description.
func (p process) MSet(values ...interface{}) *StatusCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = "mset"
	args = appendArgs(args, values)
	cmd := NewStatusCmd(args...)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

// Set Redis `SET key value [expiration]` command.
// Use expiration for `SETEx`-like behavior.
//
// Zero expiration means the key has no expiration time.
// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
// otherwise you will receive an error: (error) ERR syntax error.
func (p process) Set(key string, value interface{}) *StatusCmd {
	args := make([]interface{}, 3, 5)
	args[0] = "set"
	args[1] = key
	args[2] = value

	cmd := NewStatusCmd(args...)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func (p process) StrLen(key string) *IntCmd {
	cmd := NewIntCmd("strlen", key)
	cmd.SetOp(konata_client.Write)
	_ = p(cmd)
	return cmd
}

func appendArgs(dst, src []interface{}) []interface{} {
	dst = append(dst, src...)
	return dst
}
