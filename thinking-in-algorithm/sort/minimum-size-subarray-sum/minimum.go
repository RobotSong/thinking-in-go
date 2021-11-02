package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	l, r := 1, len(nums)+1

	for l < r {
		m := l + ((r - l) >> 1)
		mov := getC(nums, m, target)
		if mov < 0 {
			l = m + 1
		} else {
			r = m
		}
	}

	if l > len(nums) {
		return 0
	}

	return l
}

func getC(nums []int, l int, target int) int {
	sum := 0
	// 相当于 一个长度为 l 的滑动窗口
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 长度还没到 l 时
		if i < l-1 {
			continue
		}

		if sum >= target {
			return 0
		}
		// 滑动窗口的实现
		sum -= nums[i-(l-1)]
	}

	return -1
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 3
	l := minSubArrayLen(s, nums)
	fmt.Println(l)
}
