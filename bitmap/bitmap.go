package bitmap

import "fmt"

// bitmap struct
type BitMap struct {
	bits    []byte // 保存实际的 bit 数据
	bitSize uint64 // 指示该 Bitmap 能够存储的最大数字
}

// 初始化bitmap
func NewBitMap(bitSize uint64) *BitMap {
	bits := make([]byte, (bitSize>>3)+1)
	return &BitMap{bits: bits, bitSize: bitSize}
}

// 添加数据到位图
func (b *BitMap) Add(num uint64) {
	index := num / 8 // 先找到整个slice中的索引位置
	pos := num % 8   // 找到在这个byte中所在位的位置

	b.bits[index] |= 1 << pos
}

// 删除对应位图数据
func (b *BitMap) Delete(num uint64) {
	index := num / 8 // 先找到整个slice中的索引位置
	pos := num % 8   // 找到在这个byte中所在位的位置

	b.bits[index] &^= 1 << pos
}

// 检查是否在位图中
func (b *BitMap) Has(num uint64) bool {
	index := num / 8 // 先找到整个slice中的索引位置
	pos := num % 8   // 找到在这个byte中所在位的位置
	if index > uint64(len(b.bits)) {
		return false
	}

	return b.bits[index]&(1<<pos) != 0
}

//
func (b *BitMap) String() string {
	return fmt.Sprint(b.bits)
}
