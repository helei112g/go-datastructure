// 整数数组实现大顶堆
package heap

type MyHeap struct {
	heap []int32
}

// 初始化一个堆
// 从后往前堆化；又叶子节点只能跟自己比较，因此从 hLen/2 - 1 节点开始处理
func Init(nums []int32) *MyHeap {
	hLen := len(nums)
	for i := hLen/2 - 1; i >= 0; i-- {
		down(nums, hLen, i)
	}

	h := &MyHeap{heap: nums}
	return h
}

// 添加一个元素
// 先放到数组最后，逐级进行比较
func (h *MyHeap) Push(num int32) {
	h.heap = append(h.heap, num)

	up(h.heap, len(h.heap)-1)
}

// 弹出一个元素
func (h *MyHeap) Pop() int32 {
	if len(h.heap) == 0 {
		return 0
	}

	top := h.heap[0] // 弹出堆顶元素

	n := len(h.heap) - 1
	// 将最后一个元素放到顶部，然后重新堆化
	h.heap[0] = h.heap[n]
	h.heap = h.heap[:n]
	down(h.heap, n, 0)

	return top
}

// 构建堆
// 从下往上堆化：与父子节点比，大就交换
func up(nums []int32, pos int) {
	for {
		parent := (pos - 1) / 2 // parent pos
		if parent == pos || nums[pos] <= nums[parent] {
			break
		}

		// swap
		nums[pos], nums[parent] = nums[parent], nums[pos]

		pos = parent
	}
}

// 堆化的过程，找出 pos 进行交换的位置
// 与左右子节点进行比较，由于下标是从0开始，左节点是：2*pos+1; 右节点是：2*pos+2
// 构建堆，从上往下
func down(nums []int32, hLen, pos int) {
	for {
		l := 2*pos + 1          // left child
		if l >= hLen || l < 0 { // out of range
			break
		}
		j := l     // default left pos
		r := l + 1 // right child
		if r < hLen && nums[l] < nums[r] {
			j = r // 2*pos + 2, set right pos
		}

		// 父节点与最大的子节点进行比较，如果父节点已经最大，终止
		if nums[pos] >= nums[j] { // don't need swap
			break
		}
		nums[pos], nums[j] = nums[j], nums[pos]

		pos = j
	}
}
