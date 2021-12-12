package main

import "fmt"

func searchInsert(nums []int, target int) int {
	// 开闭区间
	l, r := 0, len(nums)

	for l < r {
		m := l + ((r - l) >> 1)
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	ans := searchInsert(nums, target)
	fmt.Println(ans)
}
