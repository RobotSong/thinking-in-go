package fibo

func RecursiveFibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
}
