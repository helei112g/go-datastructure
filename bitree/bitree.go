package bitree

import "fmt"

// k 要查找的值
// f 指向当前t的双亲节点
func (t *TNode) SearchBST(k int64, f *TNode) (ret *TNode, ok bool) {
	if t == nil {
		ret = f
		return ret, false
	} else if t.data == k {
		ret = t
		return ret, true
	} else if t.data < k {
		// 查找值，大于该节点，向右查找
		return t.rChild.SearchBST(k, t)
	} else {
		return t.lChild.SearchBST(k, t)
	}
}

// 向二叉排序树添加元素
func (t *TNode) InsertBST(k int64) *TNode {
	if ret, ok := t.SearchBST(k, t); !ok {
		// 未找到插入
		n := &TNode{
			data:   k,
			lChild: nil,
			rChild: nil,
		}

		if ret == nil {
			return n
		} else if ret.data > k {
			ret.lChild = n
		} else {
			ret.rChild = n
		}

	}

	return t
}

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
