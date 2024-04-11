package Konata_client

type SetCmdable interface {
	// set
	SAdd(key string, members ...interface{}) *IntCmd
	SCard(key string) *IntCmd
	SDiff(keys ...string) *StringSliceCmd
	SInter(keys ...string) *StringSliceCmd
	SRem(key string, members ...interface{}) *IntCmd
	SUnion(keys ...string) *StringSliceCmd
}

func (p process) SAdd(key string, members ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(members))
	args[0] = "sadd"
	args[1] = key
	args = appendArgs(args, members)
	cmd := NewIntCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) SCard(key string) *IntCmd {
	cmd := NewIntCmd("scard", key)
	_ = p(cmd)
	return cmd
}

func (p process) SDiff(keys ...string) *StringSliceCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = "sdiff"
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewStringSliceCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) SRem(key string, members ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(members))
	args[0] = "srem"
	args[1] = key
	args = appendArgs(args, members)
	cmd := NewIntCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) SInter(keys ...string) *StringSliceCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = "sinter"
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewStringSliceCmd(args...)
	_ = p(cmd)
	return cmd
}

func (p process) SUnion(keys ...string) *StringSliceCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = "sunion"
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewStringSliceCmd(args...)
	_ = p(cmd)
	return cmd
}
