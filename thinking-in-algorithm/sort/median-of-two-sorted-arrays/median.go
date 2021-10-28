package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 长度
	len1 := len(nums1)
	len2 := len(nums2)
	lenTotal := len1 + len2
	// 如果两个数组的总长度为0 ， 那么肯定没有中位数
	if lenTotal == 0 {
		return 0
	}
	// 总长度为偶数的情况：
	// 如果有4个数，那么当扔掉 1 个之后
	// 接下来需要合并的两个数排 [2, 3]
	// 总长度为奇数的情况：
	// 如果有5个数，那么当合并掉 2 个数之后
	// 接下来需要的那排 [3] 位的就是中位数
	// 所以这里的 k 表示：要扔掉的个数
	i := 0
	j := 0
	k := (lenTotal - 1) >> 1
	for k > 0 {
		// 我们需要比较 nums1[p] 和 nums2[p]
		// 只不过当数组的起始位置是 i 和 j  的时候。
		// 比较的元素就变成 nums1[i + p] 和 nums2[j + p]
		p := (k - 1) >> 1
		// 这时直接比较 nums1[i + p] 和 nums2[j + p] 来决定谁可以被扔掉
		// 注意这里扔掉的时候，只需要前移 p + 1 即可
		if j+p >= len2 || i+p < len1 && nums1[i+p] < nums2[j+p] {
			i += p + 1
		} else {
			j += p + 1
		}

		k -= p + 1
	}

	// 取出前面的数
	front := 0.0
	if j >= len2 || i < len1 && nums1[i] < nums2[j] {
		front = float64(nums1[i])
		i++
	} else {
		front = float64(nums2[j])
		j++
	}

	// 如果总长度是奇数，那么这个时候，front 就是我们要找的中位数
	if (lenTotal & 1) == 1 {
		return front
	}

	// 如果总长度是偶数，那么需要再取一个数，求平均值
	back := 0.0
	if j >= len2 || i < len1 && nums1[i] < nums2[j] {
		back = float64(nums1[i])
	} else {
		back = float64(nums2[j])
	}
	return (front + back) / 2
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}

	mid := findMedianSortedArrays(nums1, nums2)
	fmt.Println(mid)
}
