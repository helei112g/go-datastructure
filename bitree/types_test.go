package bitree

import (
	"reflect"
	"testing"
)

func TestNewBiTree(t *testing.T) {
	type args struct {
		in []int64
	}
	tests := []struct {
		name string
		args args
		want *TNode
	}{
		{
			name: "init",
			args: args{in: []int64{62, 88, 58, 47, 35, 73, 51, 99, 37, 93}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBiTree(tt.args.in...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBiTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
