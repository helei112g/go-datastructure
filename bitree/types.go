package bitree

import "fmt"

type TNode struct {
	data           int64
	lChild, rChild *TNode
}

func NewBiTree(in ...int64) *TNode {
	var root *TNode
	for _, val := range in {
		root = root.InsertBST(val)
	}

	fmt.Println(root)
	return root
}
