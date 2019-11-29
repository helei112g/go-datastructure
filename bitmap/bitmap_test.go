package bitmap

import (
	"testing"
)

func TestNewBitMap(t *testing.T) {
	var bitSize uint64
	bitSize = 9999999999

	b := NewBitMap(bitSize)
	b.Add(9999999999)
	b.Add(0)
	b.Add(199999)

	if flag := b.Has(9999999999); flag != true {
		t.Errorf("b.Has = %v, want = %v", flag, true)
	}

	if flag := b.Has(0); flag != true {
		t.Errorf("b.Has = %v, want = %v", flag, true)
	}

	if flag := b.Has(199999); flag != true {
		t.Errorf("b.Has = %v, want = %v", flag, true)
	}

	// 测试删除
	b.Delete(199999)
	if flag := b.Has(199999); flag == true {
		t.Errorf("b.Has = %v, want = %v", flag, true)
	}
}
