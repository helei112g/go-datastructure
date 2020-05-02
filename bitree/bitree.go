package bitree

import "fmt"

// 前序遍历
func (t *TNode) PreOrderTraverse() {
	if t == nil {
		return
	}

	fmt.Println(t.data)
	t.lChild.PreOrderTraverse()
	t.rChild.PreOrderTraverse()
}

func (t *TNode) MiddleOrderTraverse() {
	if t == nil {
		return
	}

	t.lChild.MiddleOrderTraverse()
	fmt.Println(t.data)
	t.rChild.MiddleOrderTraverse()
}

func (t *TNode) PostOrderTraverse() {
	if t == nil {
		return
	}

	t.lChild.PostOrderTraverse()
	t.rChild.PostOrderTraverse()
	fmt.Println(t.data)
}
