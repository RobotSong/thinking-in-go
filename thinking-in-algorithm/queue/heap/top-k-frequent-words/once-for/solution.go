package main

import "fmt"

type Heap struct {
	a []node
	// 字符串所在下标的位置
	im map[string]int
	n  int
}

type node struct {
	w string
	c int
}

func (h *Heap) Size() int {
	return h.n
}

func (h *Heap) less(a, b node) bool {
	return a.c < b.c || a.c == b.c && a.w > b.w
}

// 下沉
func (h *Heap) sink(i int) {
	t := h.a[i]
	j := 0
	// 子结点 2 * i + 1
	for j = i<<1 + 1; j < h.n; j = i<<1 + 1 {
		// 如果存在右结点, 并且右结点更小
		if j < h.n-1 && h.less(h.a[j+1], h.a[j]) {
			j++
		}
		// 如果子结点比 t 更小
		if h.less(h.a[j], t) {
			h.im[h.a[j].w] = i
			// 那么子结点的位置向前排
			h.a[i] = h.a[j]
			i = j
		} else {
			// 找到了 t 的位置
			break
		}
	}

	h.a[i] = t
	h.im[t.w] = i
}

// 上浮
func (h *Heap) swim(i int) {
	t := h.a[i]
	var par int
	for i > 0 {
		// 父结点 (i - 1) / 2
		par = (i - 1) >> 1
		// 如果父结点比 t 大,
		if h.less(t, h.a[par]) {
			h.im[h.a[par].w] = i
			// 那么父结点位置向下移动
			h.a[i] = h.a[par]
			i = par
			//

		} else {
			// 找到了 t 的位置
			break
		}
	}
	h.a[i] = t
	h.im[t.w] = i
}

func (h *Heap) Push(w string) {
	i, has := h.im[w]
	if has {
		// 如果已经存在, c + 1, 并且从 i 开始下沉
		n := h.a[i]
		n.c++
		h.a[i] = n
		h.sink(i)
	} else {

		// 如果不存在, 则放在最后 并上浮
		n := node{
			w: w,
			c: 1,
		}
		if len(h.a) == h.n {
			h.a = append(h.a, n)
		} else {
			h.a[h.n] = n
		}

		h.im[w] = h.n
		h.n++
		h.swim(h.n - 1)
	}
}

func (h *Heap) pop() node {
	// 取出最小值
	n := h.a[0]
	// 从最后一位赋值到 首位
	h.a[0] = h.a[h.n-1]
	h.n--
	// 并且从 0 开始下沉
	h.sink(0)
	return n
}

func Construct() Heap {
	return Heap{
		im: make(map[string]int),
		n:  0,
	}
}

func topKFrequent(words []string, k int) []string {

	h := Construct()

	for _, w := range words {
		h.Push(w)
	}

	for h.Size() > k {
		h.pop()
	}

	var result = make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = h.pop().w
	}

	return result
}

func main() {
	words := []string{"glarko", "zlfiwwb", "nsfspyox", "pwqvwmlgri", "qggx", "qrkgmliewc", "zskaqzwo", "zskaqzwo", "ijy", "htpvnmozay", "jqrlad", "ccjel", "qrkgmliewc", "qkjzgws", "fqizrrnmif", "jqrlad", "nbuorw", "qrkgmliewc", "htpvnmozay", "nftk", "glarko", "hdemkfr", "axyak", "hdemkfr", "nsfspyox", "nsfspyox", "qrkgmliewc", "nftk", "nftk", "ccjel", "qrkgmliewc", "ocgjsu", "ijy", "glarko", "nbuorw", "nsfspyox", "qkjzgws", "qkjzgws", "fqizrrnmif", "pwqvwmlgri", "nftk", "qrkgmliewc", "jqrlad", "nftk", "zskaqzwo", "glarko", "nsfspyox", "zlfiwwb", "hwlvqgkdbo", "htpvnmozay", "nsfspyox", "zskaqzwo", "htpvnmozay", "zskaqzwo", "nbuorw", "qkjzgws", "zlfiwwb", "pwqvwmlgri", "zskaqzwo", "qengse", "glarko", "qkjzgws", "pwqvwmlgri", "fqizrrnmif", "nbuorw", "nftk", "ijy", "hdemkfr", "nftk", "qkjzgws", "jqrlad", "nftk", "ccjel", "qggx", "ijy", "qengse", "nftk", "htpvnmozay", "qengse", "eonrg", "qengse", "fqizrrnmif", "hwlvqgkdbo", "qengse", "qengse", "qggx", "qkjzgws", "qggx", "pwqvwmlgri", "htpvnmozay", "qrkgmliewc", "qengse", "fqizrrnmif", "qkjzgws", "qengse", "nftk", "htpvnmozay", "qggx", "zlfiwwb", "bwp", "ocgjsu", "qrkgmliewc", "ccjel", "hdemkfr", "nsfspyox", "hdemkfr", "qggx", "zlfiwwb", "nsfspyox", "ijy", "qkjzgws", "fqizrrnmif", "qkjzgws", "qrkgmliewc", "glarko", "hdemkfr", "pwqvwmlgri"}
	k := 14

	// 预期结果
	// ["nftk","qkjzgws","qrkgmliewc","nsfspyox","qengse","htpvnmozay","fqizrrnmif","glarko","hdemkfr","pwqvwmlgri","qggx","zskaqzwo","ijy","zlfiwwb"]
	result := topKFrequent(words, k)
	fmt.Println(result)
}
