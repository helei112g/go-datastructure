package bitree

import (
	"testing"
)

func getTree() *TNode {
	return NewBiTree(62, 88, 58, 47, 35, 73, 51, 99, 37, 93)
}

func TestTNode_PreOrderTraverse(t1 *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test01",
		},
	}

	tree := getTree()
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tree.PreOrderTraverse()
		})
	}
}

func TestTNode_MiddleOrderTraverse(t1 *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test01",
		},
	}

	tree := getTree()
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tree.MiddleOrderTraverse()
		})
	}
}

func TestTNode_PostOrderTraverse(t1 *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test01",
		},
	}
	tree := getTree()
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tree.PostOrderTraverse()
		})
	}
}
