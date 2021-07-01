package fib

import (
	"math/rand"
	"testing"
)

func TestFibonacci(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"FibonacciV2", args{1}, 1},
		{"FibonacciV2", args{2}, 1},
		{"FibonacciV2", args{3}, 2},
		{"FibonacciV2", args{4}, 3},
		{"FibonacciV2", args{5}, 5},
		{"FibonacciV2", args{6}, 8},
		{"FibonacciV2", args{7}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.x); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibonacciV2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciV2(rand.Int() % 20000)
	}
}
func BenchmarkFibonacciV1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciV1(rand.Int() % 20000)
	}
}
func BenchmarkFibonacci(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Fibonacci(rand.Int() % 50)
	}
}
