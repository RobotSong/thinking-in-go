package main

import (
	"fmt"
)

// 堆中的元素
type item struct {
	word  string
	count int
}

// Heap 根据 item.count 为小堆 如果 count 相等, 根据 word 大的排前面
type Heap struct {
	arr []item
	n   int
}

func Construct(k int) Heap {
	return Heap{
		arr: make([]item, k+1, k+1),
		n:   0,
	}
}

// Size 当前堆中的数据大小
func (h *Heap) Size() int {
	return h.n
}

// sink 下沉元素
func (h *Heap) sink(i int) {

	var t = h.arr[i]
	var j = 0
	// 是否还存在子结点
	for j = i<<1 + 1; j < h.n; j = i<<1 + 1 {
		// 如果存在右结点， 并且右结点的 count 更小
		if j < h.n-1 && h.less(h.arr[j+1], h.arr[j]) {
			j++
		}
		// 如果子结点的 count 比 t 的小
		// 那么 t 的位置要往后排
		if h.less(h.arr[j], t) {
			h.arr[i] = h.arr[j]
			i = j
		} else {
			// 找到了 t 的位置
			break
		}

	}
	h.arr[i] = t
}

// swim 上浮元素
func (h Heap) swim(i int) {

	t := h.arr[i]
	// 父结点的 位置
	var par int
	for i > 0 {
		// (i - 1) / 2
		par = (i - 1) >> 1
		// 如果父结点 比子节点大
		if h.less(t, h.arr[par]) {
			// 那么向下移动父结点的值
			h.arr[i] = h.arr[par]
			i = par
		} else {
			// 找到了 t 的位置
			break
		}
	}

	h.arr[i] = t
}

// pop 取出 item.count 最小的元素
func (h *Heap) pop() item {
	// pop 取出 最小值后，无法获取到下一个最小值
	ret := h.arr[0]
	// 取出后，将最后一位赋值到 0 ，然后下沉该元素
	h.arr[0] = h.arr[h.n-1]
	h.arr[h.n-1] = item{}
	h.n--
	h.sink(0)
	return ret
}

// push 放入最新的值
func (h *Heap) push(val item) {
	h.arr[h.n] = val
	h.n++
	// 上浮
	h.swim(h.n - 1)
}

// getFirst 获取第一个元素
func (h *Heap) getFirst() item {
	return h.arr[0]
}

func (h Heap) less(a, b item) bool {
	return a.count < b.count || a.count == b.count && a.word > b.word
}

func topKFrequent(words []string, k int) []string {
	var countMap = make(map[string]int)

	for _, word := range words {
		count := countMap[word]
		countMap[word] = count + 1
	}

	h := Construct(k)
	for word, count := range countMap {
		h.push(item{
			word:  word,
			count: count,
		})
		if h.Size() > k {
			h.pop()
		}

	}

	var result = make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = h.pop().word
	}
	for _, word := range result {
		fmt.Printf("%s:%d\n", word, countMap[word])
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
