package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {

	h := Construct(100)
	h.Push(3)
	h.Push(8)
	h.Push(6)
	h.Push(2)
	h.Push(5)
	h.Push(1)
	h.Push(2)
	h.Push(2)
	h.Push(1)
	h.Push(1)
	h.Push(8)
	h.Push(9)
	h.Push(11)

	ret := h.Pop()
	if ret != 8 {
		t.Errorf("want : 8, got: %d", ret)
	}
	ret = h.Pop()
	if ret != 6 {
		t.Errorf("want : 6, got: %d", ret)
	}
	//ret = h.Size()
	//if ret != 3 {
	//	t.Errorf("want : 3, got: %d", ret)
	//}
	ret = h.Pop()
	if ret != 5 {
		t.Errorf("want : 5, got: %d", ret)
	}
	h.Pop()
	h.Pop()
	ret = h.Pop()
	if ret != -1 {
		t.Errorf("want : -1, got: %d", ret)
	}
}
