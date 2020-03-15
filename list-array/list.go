package list_array

const MaxLen = 20

type ElemType int

type SqList struct {
	// 存储元素
	list [MaxLen]ElemType
	// 元素个数
	len int
}

// 定义操作
type List interface {
	// 检查是否为空
	IsEmpty() bool
	// 清空链表
	ClearList() bool
	// 获取位置是i的元素
	GetElem(i int) (e ElemType, ok bool)
	// 查找等于e的元素位置
	GetElemLocation(e ElemType) (i int, ok bool)
	// 获取长度
	GetLen() int
	// 插入元素
	Insert(e ElemType, i int) bool
	// 删除元素
	Delete(i int) (e ElemType, ok bool)
}

// 初始化链表
func InitList() *SqList {
	return &SqList{
		list: [20]ElemType{},
		len:  0,
	}
}

// 是否为空
func (l *SqList) IsEmpty() bool {
	if l.len == 0 {
		return true
	}
	return false
}

// 清空链表
// 清空时不处理原有元素，使用时直接覆盖
func (l *SqList) ClearList() bool {
	l.len = 0
	return true
}

// 获取元素
func (l *SqList) GetElem(i int) (e ElemType, ok bool) {
	if i < 1 || l.len < i || l.len == 0 {
		return 0, false
	}

	return l.list[i-1], true
}

// 获取元素的位置
func (l *SqList) GetElemLocation(e ElemType) (i int, ok bool) {
	if l.len == 0 {
		return
	}

	for index, val := range l.list {
		if e == val {
			return index + 1, true
		}
	}

	return
}

// 获取链表长度
func (l *SqList) GetLen() int {
	return l.len
}

// 插入元素
func (l *SqList) Insert(e ElemType, i int) bool {
	if l.len == MaxLen {
		return false
	}
	if i < 1 || i > l.len+1 {
		return false
	}

	// 数据在中间位置，需要先移动
	if i < l.len {
		for k := l.len - 1; k >= i-1; k++ {
			l.list[k+1] = l.list[k]
		}
	}

	l.list[i-1] = e
	l.len++
	return true
}

// 删除操作
func (l *SqList) Delete(i int) (e ElemType, ok bool) {
	if l.len == 0 {
		return
	}
	if i < 1 || i > l.len {
		return
	}

	e = l.list[i-1]

	// 删除后移动元素位置
	for k := i; k < l.len; k++ {
		l.list[k-1] = l.list[k]
	}
	l.len--
	return e, true
}
