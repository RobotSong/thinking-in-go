package fibo

import (
	"testing"
)

func BenchmarkRecursiveFibonacci_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibonacci(10)
	}
}

func BenchmarkRecursiveFibonacci_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibonacci(20)
	}
}

func BenchmarkSequentialFibonacci_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialFibonacci(10)
	}
}

func BenchmarkSequentialFibonacci_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialFibonacci(20)
	}
}

func TestRecursiveFibonacci(t *testing.T) {

	data := []struct {
		n    uint
		want uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{11, 89},
		{42, 267914296},
	}

	for _, d := range data {
		if got := RecursiveFibonacci(d.n); got != d.want {
			t.Errorf("got: %d, want: %d\n", got, d.want)
		}
	}
}

func TestSequentialFibonacci(t *testing.T) {
	data := []struct {
		n    uint
		want uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{11, 89},
		{42, 267914296},
	}

	for _, d := range data {
		if got := SequentialFibonacci(d.n); got != d.want {
			t.Errorf("got: %d, want: %d\n", got, d.want)
		}
	}
}
