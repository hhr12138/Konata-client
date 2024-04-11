package Konata_client

type HashCmdable interface {
	// hash
	HSet(key, field string, value interface{}) *BoolCmd
	HGet(key, field string) *StringCmd
	HDel(key string, fields ...string) *IntCmd
	HExists(key, field string) *BoolCmd
	HGetAll(key string) *StringStringMapCmd
	HLen(key string) *IntCmd
	HIncrBy(key, field string, incr int64) *IntCmd
	HIncrByFloat(key, field string, incr float64) *FloatCmd
	HDecrBy(key, field string, incr int64) *IntCmd
	HDecrByFloat(key, field string, incr float64) *FloatCmd
	HVals(key string) *StringSliceCmd
}

func (p process) HDel(key string, fields ...string) *IntCmd {
	args := make([]interface{}, 2+len(fields))
	args[0] = "hdel"
	args[1] = key
	for i, field := range fields {
		args[2+i] = field
	}
	cmd := NewIntCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) HExists(key, field string) *BoolCmd {
	cmd := NewBoolCmd("hexists", key, field)
	_ = p(cmd)
	return cmd
}

func (p process) HGet(key, field string) *StringCmd {
	cmd := NewStringCmd("hget", key, field)
	_ = p(cmd)
	return cmd
}

func (p process) HSet(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(values))
	args[0] = "hset"
	args[1] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) HGetAll(key string) *StringStringMapCmd {
	cmd := NewStringStringMapCmd("hgetall", key)
	_ = p(cmd)
	return cmd
}

func (p process) HLen(key string) *IntCmd {
	cmd := NewIntCmd("hlen", key)
	_ = p(cmd)
	return cmd
}

func (p process) HIncrBy(key, field string, incr int64) *IntCmd {
	cmd := NewIntCmd("hincrby", key, field, incr)
	_ = p(cmd)
	return cmd
}

func (p process) HIncrByFloat(key, field string, incr float64) *FloatCmd {
	cmd := NewFloatCmd("hincrbyfloat", key, field, incr)
	_ = p(cmd)
	return cmd
}

func (p process) HDecrBy(key, field string, incr int64) *IntCmd {
	cmd := NewIntCmd("hdecrby", key, field, incr)
	_ = p(cmd)
	return cmd
}

func (p process) HDecrByFloat(key, field string, incr float64) *FloatCmd {
	cmd := NewFloatCmd("hdecrbyfloat", key, field, incr)
	_ = p(cmd)
	return cmd
}

func (p process) HVals(key string) *StringSliceCmd {
	cmd := NewStringSliceCmd("hvals", key)
	_ = p(cmd)
	return cmd
}
