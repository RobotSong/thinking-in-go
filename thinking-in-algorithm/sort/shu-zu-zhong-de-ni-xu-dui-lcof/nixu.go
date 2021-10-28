package main

import "fmt"

func reversePairs(nums []int) int {
	t := make([]int, len(nums))
	return msort(nums, 0, len(nums), t)
}

// 合并排序数组
func msort(a []int, b, e int, t []int) int {
	result := 0
	// 如果区别为空，或者区间只有一个元素
	if b >= e || b+1 >= e {
		return result
	}
	// 查找数组的中间下标
	var m = b + ((e - b) >> 1)

	result += msort(a, b, m, t)
	result += msort(a, m, e, t)

	i := b
	j := m
	to := b
	// 合并排序
	for i < m || j < e {
		if j >= e || i < m && a[i] <= a[j] {
			t[to] = a[i]
			to++
			i++
			result += j - m
		} else {
			t[to] = a[j]
			to++
			j++
		}
	}

	for i = b; i < e; i++ {
		a[i] = t[i]
	}
	return result
}

func main() {
	nums := []int{7, 5, 6, 4}
	result := reversePairs(nums)
	fmt.Println(result)
}
