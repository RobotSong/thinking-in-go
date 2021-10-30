package main

import "fmt"

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

func sortColors(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	// 三路划分
	l := 0
	i := 0
	r := len(nums) - 1

	for i <= r {
		if nums[i] == 0 {
			swap(nums, l, i)
			l++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			swap(nums, r, i)
			r--
		}
	}

}

func main() {
	nums := []int{0, 1, 1, 2, 0, 0}
	sortColors(nums)
	fmt.Println(nums)
}
