package Konata_client

type ListCmdable interface {
	// list,虽然命令对齐redis，但这里的list实质上是数组
	LIndex(key string, index int64) *StringCmd
	LPush(key string, values ...interface{}) *IntCmd
	LPushX(key string, value interface{}) *IntCmd
	RPush(key string, values ...interface{}) *IntCmd
	RPushX(key string, value interface{}) *IntCmd
	LPop(key string) *StringCmd
	LPopCount(key string, count int64) *StringSliceCmd
	RPop(key string) *StringCmd
	RPopCount(key string, count int64) *StringSliceCmd
	LRange(key string, start, stop int64) *StringSliceCmd
	LRem(key string, count int64, value interface{}) *IntCmd
	LSet(key string, index int64, value interface{}) *StatusCmd
	LInsert(key, op string, pivot, value interface{}) *IntCmd
}

func (p process) LIndex(key string, index int64) *StringCmd {
	cmd := NewStringCmd("lindex", key, index)
	_ = p(cmd)
	return cmd
}

func (p process) LPush(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(values))
	args[0] = "lpush"
	args[1] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) LPushX(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(values))
	args[0] = "lpushx"
	args[1] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) RPush(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(values))
	args[0] = "rpush"
	args[1] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) RPushX(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(values))
	args[0] = "rpushx"
	args[1] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) LPop(key string) *StringCmd {
	cmd := NewStringCmd("lpop", key)
	_ = p(cmd)
	return cmd
}

func (p process) RPopCount(key string, count int) *StringSliceCmd {
	cmd := NewStringSliceCmd("rpop", key, count)
	_ = p(cmd)
	return cmd
}

func (p process) RPop(key string) *StringCmd {
	cmd := NewStringCmd("rpop", key)
	_ = p(cmd)
	return cmd
}

func (p process) LRange(key string, start, stop int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(
		"lrange",
		key,
		start,
		stop,
	)
	_ = p(cmd)
	return cmd
}

func (p process) LRem(key string, count int64, value interface{}) *IntCmd {
	cmd := NewIntCmd("lrem", key, count, value)
	_ = p(cmd)
	return cmd
}

func (p process) LSet(key string, index int64, value interface{}) *StatusCmd {
	cmd := NewStatusCmd("lset", key, index, value)
	_ = p(cmd)
	return cmd
}

func (p process) LInsert(key, op string, pivot, value interface{}) *IntCmd {
	cmd := NewIntCmd("linsert", key, op, pivot, value)
	_ = p(cmd)
	return cmd
}
