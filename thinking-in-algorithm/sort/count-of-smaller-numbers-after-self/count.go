package main

import "fmt"

func countSmaller(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return make([]int, 0)
	}

	n := len(nums)
	ans, idx, t := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	msort(nums, 0, n, ans, idx, t)
	return ans
}

func msort(a []int, b, e int, ans, idx, t []int) {
	if b >= e || b+1 >= e {
		return
	}

	m := b + ((e - b) >> 1)
	msort(a, b, m, ans, idx, t)
	msort(a, m, e, ans, idx, t)

	i := b
	j := m
	to := b

	for i < m || j < e {
		if j >= e || i < m && a[idx[i]] <= a[idx[j]] {
			ans[idx[i]] += j - m
			t[to] = idx[i]
			to++
			i++
		} else {
			t[to] = idx[j]
			to++
			j++
		}
	}

	for i = b; i < e; i++ {
		idx[i] = t[i]
	}
}

func main() {
	nums := []int{5, 2, 6, 1}
	ans := countSmaller(nums)
	fmt.Println(ans)
}
