package main

import "fmt"

func lowerBound(nums []int, target int) int {
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

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + ((r - l) >> 1)
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return ans
	}

	l := lowerBound(nums, target)
	r := upperBound(nums, target)
	if l == r {
		return ans
	}
	ans[0] = l
	ans[1] = r - 1

	return ans
}

func main() {
	nums := []int{0}
	target := 0

	ans := searchRange(nums, target)
	fmt.Println(ans)
}
