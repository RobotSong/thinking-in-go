package main

import "fmt"

// 获取临时数组 C 的值
func getC(arr []int, i int) int {
	// 为升序的山峰
	if arr[i-1] < arr[i] && arr[i] < arr[i+1] {
		return -1
	} else if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
		return 0
	} else if arr[i-1] > arr[i] && arr[i] > arr[i+1] {
		// 为降序的山峰
		return 1
	}

	return 1
}

func peakIndexInMountainArray(arr []int) int {
	if arr == nil || len(arr) < 3 {
		return -1
	}
	// 因为山峰至少需要三个元素组成, 并且判断时, 需要 -1 所以 l 从 1开始
	l, r := 1, len(arr)-1
	// 使用 lowerBound 来查找为 0 (峰顶元素) 的值
	for l < r {
		m := l + ((r - l) >> 1)
		if getC(arr, m) < 0 {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func main() {
	nums := []int{0, 1, 0}
	i := peakIndexInMountainArray(nums)
	fmt.Println(i)
}
