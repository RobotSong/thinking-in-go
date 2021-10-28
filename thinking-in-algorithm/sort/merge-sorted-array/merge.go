package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 从大到小
	tail := m + n - 1
	i := m - 1
	j := n - 1
	for i >= 0 || j >= 0 {

		if j < 0 || i >= 0 && nums1[i] >= nums2[j] {
			nums1[tail] = nums1[i]
			tail--
			i--
		} else {
			nums1[tail] = nums2[j]
			tail--
			j--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

}
