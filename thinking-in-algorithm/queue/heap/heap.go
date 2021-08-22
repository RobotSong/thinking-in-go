package heap

type Heap struct {
	a []int
	n int
}

func Construct(k int) Heap {
	return Heap{
		a: make([]int, k, k),
		n: 0,
	}
}

// 下沉
func (h *Heap) sink(i int) {
	// 下沉的值
	var t int = h.a[i]
	// 还存在子节点, 每次循环重新获取左子节点的下标
	// i 结点的左结点为 2 * i + 1, 右结点为 2 * i + 2
	for j := i<<1 + 1; j < h.n; j = i<<1 + 1 {
		// 如果还存在右子结点 ， 并且右子结点 比左结点大
		if j < h.n-1 && h.a[j] < h.a[j+1] {
			j++
		}
		// 如果 子结点的值 比 t大
		// 那么 t 的位置还需要往后排
		if h.a[j] > t {
			h.a[i] = h.a[j]
			i = j
		} else {
			break
		}
	}

	// 将 t 放在找到的位置那里
	h.a[i] = t
}

// 上浮
func (h *Heap) swim(i int) {
	t := h.a[i]
	// 父结点的位置
	var par int
	for i > 0 {
		// i 结点的父结点 (i - 1) / 2
		par = (i - 1) >> 1
		// 如果父结点 比 t小
		if h.a[par] < t {
			// 那么向下移动父结点的值
			h.a[i] = h.a[par]
			i = par
		} else {
			// 找到了 t 的位置
			break
		}
	}
	// 将 t 放在找到的位置那里
	h.a[i] = t
}

func (h *Heap) Size() int {
	return h.n
}

func (h *Heap) Push(val int) {
	// push 是先把元素追加到数组尾巴上，然后再执行上浮操作
	h.a[h.n] = val
	h.n++
	h.swim(h.n - 1)
}

func (h *Heap) Pop() int {
	if h.n == 0 {
		return -1
	}
	// 先取出 a[0] 作为返回值
	ret := h.a[0]
	// 将 a[n - 1] 存放至 a[0]
	h.a[0] = h.a[h.n-1]
	h.n--
	// 并将 a[0] 进行下沉操作
	h.sink(0)
	return ret
}
