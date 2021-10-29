// 使用快速排序
package main

import "fmt"

// swap 交换数组元素
func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

func quick(nums []int, b, e int) {
	// 如果区间没有元素，或者只有一个元素
	if b >= e || b+1 >= e {
		return
	}

	// 找到中位数
	m := b + ((e - b) >> 1)
	x := nums[m]

	l := b
	i := b
	r := e - 1

	for i <= r {
		if nums[i] < x {
			// 如果是小于 x 的数，则放到左区间，左区间大小变大
			swap(nums, l, i)
			l++
			i++
		} else if nums[i] == x {
			// 如果相等, 则中间区间大小变大
			i++
		} else {
			// 如果大于，则右区间大小变小，然后 r,i 交换， 但 i 的值不变，等待下一次的判断
			swap(nums, r, i)
			r--
		}
	}

	// 前序遍历
	quick(nums, b, l)
	quick(nums, i, e)

}

func sortArray(nums []int) []int {
	quick(nums, 0, len(nums))
	return nums
}

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	nums = sortArray(nums)
	fmt.Println(nums)
}
