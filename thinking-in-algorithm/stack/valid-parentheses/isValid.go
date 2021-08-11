package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		// 如果是右半部分，则进行出栈，判断是否匹配操作
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			// 如果是左半部分，则进行 入栈操作
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	data := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[()}]", false},
		{"{[]}", true},
	}

	for _, s := range data {
		if got := isValid(s.s); got != s.want {
			fmt.Printf("s: %q -- got: %v, want: %v", s.s, got, s.want)
		}
	}
}
