package Konata_client

import (
	"reflect"
	"testing"
)

func Test_process_HDecrBy(t *testing.T) {
	type args struct {
		key   string
		field string
		incr  int64
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
			if got := tt.p.HDecrBy(tt.args.key, tt.args.field, tt.args.incr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HDecrBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HDecrByFloat(t *testing.T) {
	type args struct {
		key   string
		field string
		incr  float64
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
			if got := tt.p.HDecrByFloat(tt.args.key, tt.args.field, tt.args.incr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HDecrByFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HDel(t *testing.T) {
	type args struct {
		key    string
		fields []string
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
			if got := tt.p.HDel(tt.args.key, tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HDel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HExists(t *testing.T) {
	type args struct {
		key   string
		field string
	}
	tests := []struct {
		name string
		p    process
		args args
		want *BoolCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.HExists(tt.args.key, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HGet(t *testing.T) {
	type args struct {
		key   string
		field string
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
			if got := tt.p.HGet(tt.args.key, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HGetAll(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    process
		args args
		want *StringStringMapCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.HGetAll(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HGetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HIncrBy(t *testing.T) {
	type args struct {
		key   string
		field string
		incr  int64
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
			if got := tt.p.HIncrBy(tt.args.key, tt.args.field, tt.args.incr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HIncrBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HIncrByFloat(t *testing.T) {
	type args struct {
		key   string
		field string
		incr  float64
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
			if got := tt.p.HIncrByFloat(tt.args.key, tt.args.field, tt.args.incr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HIncrByFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HLen(t *testing.T) {
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
			if got := tt.p.HLen(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HSet(t *testing.T) {
	type args struct {
		key    string
		values []interface{}
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
			if got := tt.p.HSet(tt.args.key, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process_HVals(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    process
		args args
		want *StringSliceCmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.HVals(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HVals() = %v, want %v", got, tt.want)
			}
		})
	}
}
