package fibo

func SequentialFibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	var n2, n1 uint = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n2+n1
	}

	return n2 + n1
}
