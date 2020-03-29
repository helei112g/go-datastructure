package heap

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		nums []int32
	}
	tests := []struct {
		name string
		args args
		want *MyHeap
	}{
		{
			name: "test0",
			args: args{nums: []int32{}},
			want: &MyHeap{heap: []int32{}},
		},
		{
			name: "test01",
			args: args{nums: []int32{7}},
			want: &MyHeap{heap: []int32{7}},
		},
		{
			name: "test02",
			args: args{nums: []int32{7, 5}},
			want: &MyHeap{heap: []int32{7, 5}},
		},
		{
			name: "test02",
			args: args{nums: []int32{7, 5, 19}},
			want: &MyHeap{heap: []int32{19, 5, 7}},
		},
		{
			name: "test02",
			args: args{nums: []int32{7, 5, 19, 8, 4, 1, 20, 13, 16}},
			want: &MyHeap{heap: []int32{20, 16, 19, 13, 4, 1, 7, 5, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		{
			name: "test0",
			want: 20,
		},
		{
			name: "test01",
			want: 19,
		},
		{
			name: "test02",
			want: 16,
		},
		{
			name: "test03",
			want: 13,
		},
		{
			name: "test04",
			want: 8,
		},
		{
			name: "test05",
			want: 7,
		},
		{
			name: "test06",
			want: 5,
		},
		{
			name: "test07",
			want: 4,
		},
		{
			name: "test08",
			want: 1,
		}, {
			name: "test09",
			want: 0,
		},
	}
	h := Init([]int32{7, 5, 19, 8, 4, 1, 20, 13, 16})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := h.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
			t.Log(h)
		})
	}
}

func TestMyHeap_Push(t *testing.T) {
	type args struct {
		num int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test0",
			args: args{num: 17},
		},
	}

	h := Init([]int32{7, 5, 19, 8, 4, 1, 20, 13, 16})
	t.Log(h)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h.Push(tt.args.num)
			t.Log(h)
		})
	}
}
