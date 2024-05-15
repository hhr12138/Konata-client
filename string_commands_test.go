package Konata_client

import (
	"fmt"
	"reflect"
	"testing"
)

var cli *DefaultClient

func init() {
	cli, _ = NewClient(nil)
}

func Test_appendArgs(t *testing.T) {
	type args struct {
		dst []interface{}
		src []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendArgs(tt.args.dst, tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_Append(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want *IntCmd
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				key:   "hi",
				value: "world",
			},
		},
		{
			name: "test1",
			args: args{
				key:   "hello",
				value: "world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cli.Append(tt.args.key, tt.args.value)
			fmt.Printf("rsp=%v\n", got.val)
		})
	}
}

func Test_process_Decr(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want *IntCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.Decr(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_DecrBy(t *testing.T) {
	type args struct {
		key       string
		decrement int64
	}
	tests := []struct {
		name string
		p    process
		args args
		want *IntCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.DecrBy(tt.args.key, tt.args.decrement); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecrBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_DecrByFloat(t *testing.T) {
	type args struct {
		key   string
		value float64
	}
	tests := []struct {
		name string
		p    process
		args args
		want *FloatCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.DecrByFloat(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecrByFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    process
		args args
		want *StringCmd
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				key: "hello",
			},
		},
		{
			name: "test0",
			args: args{
				key: "hi",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cli.Get(tt.args.key)
			fmt.Println(string(got.val))
		})
	}
}

func Test_process_GetDel(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    process
		args args
		want *StringCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.GetDel(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_GetRange(t *testing.T) {
	type args struct {
		key   string
		start int64
		end   int64
	}
	tests := []struct {
		name string
		p    process
		args args
		want *StringCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.GetRange(tt.args.key, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_GetSet(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		p    process
		args args
		want *StringCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.GetSet(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_Incr(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    process
		args args
		want *IntCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.Incr(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Incr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_IncrBy(t *testing.T) {
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name string
		p    process
		args args
		want *IntCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.IncrBy(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncrBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_IncrByFloat(t *testing.T) {
	type args struct {
		key   string
		value float64
	}
	tests := []struct {
		name string
		p    process
		args args
		want *FloatCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.IncrByFloat(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncrByFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_MGet(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		p    process
		args args
		want *SliceCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.MGet(tt.args.keys...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_MSet(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		p    process
		args args
		want *StatusCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.MSet(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_Set(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		p    process
		args args
		want *StatusCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.Set(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_StrLen(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    process
		args args
		want *IntCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cli.StrLen(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
