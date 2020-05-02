package bitree

import (
	"testing"
)

func getTree() *TNode {
	node2 := &TNode{
		data:   2,
		lChild: nil,
		rChild: nil,
	}
	node4 := &TNode{
		data:   4,
		lChild: nil,
		rChild: nil,
	}
	node5 := &TNode{
		data:   5,
		lChild: nil,
		rChild: nil,
	}

	node3 := &TNode{
		data:   3,
		lChild: node4,
		rChild: node5,
	}
	node1 := &TNode{
		data:   1,
		lChild: node2,
		rChild: nil,
	}
	root := &TNode{
		data:   99,
		lChild: node1,
		rChild: node3,
	}

	return root
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
