package main

import "fmt"

func binarySearch(nums []int, val int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	// 设定初始区间， 这里我们采用开闭原则 [l, r)
	l := 0
	r := len(nums)
	// 循环结束条件是整个区间为空区间，那么运行结束。
	// 使用的是开闭原则来表示一个区间，所以当 l <  r 的时候
	// 要查找的区间还不是一个空区间
	for l < r {
		// 中间
		m := l + ((r - l) >> 1)
		if nums[m] == val {
			return true
		} else if nums[m] < val {
			// 当中间值比目标值小时，需要把左边的部分扔掉。 即 [l, m]
			// 这个区间扔掉，由于我们采用的是开闭原则，所以新的区间需要是
			// [m + 1, r) 因需要将 l = m + 1
			l = m + 1
		} else {
			// 当中间值比目标值大，需要把右边的部分扔掉，即 [m, r) 这个区间扔掉。
			// 那么新区间变成 [l, m) 由于是开闭原则，只需要设置 r = m 即可
			r = m
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 8, 9, 11}

	b := binarySearch(nums, 8)
	fmt.Println(b)
}
