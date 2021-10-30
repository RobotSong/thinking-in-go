// 使用合并排序
package main

import "fmt"

func mergeSort(nums []int, b, e int, t []int) {
	if b >= e || b+1 >= e {
		return
	}

	m := b + ((e - b) >> 1)
	mergeSort(nums, b, m, t)
	mergeSort(nums, m, e, t)

	i := b
	j := m
	to := b
	for i < m || j < e {
		if j >= e || i < m && nums[i] <= nums[j] {
			t[to] = nums[i]
			i++
			to++
		} else {
			t[to] = nums[j]
			j++
			to++
		}
	}

	for i = b; i < e; i++ {
		nums[i] = t[i]
	}
}

func sortArray(nums []int) []int {
	t := make([]int, len(nums))
	mergeSort(nums, 0, len(nums), t)
	return nums
}

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	nums = sortArray(nums)
	fmt.Println(nums)
}
