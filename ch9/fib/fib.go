package fib

//Fibonacci 暴力递归版本
func Fibonacci(x int) int {
	if x<=0{
		return 0
	}else if x<=2{
		return 1
	}else {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
}

// FibonacciV1 顺序计算版本
func FibonacciV1(x int) int {
	if x == 0 {
		return 0
	}
	if x < 2 {
		return 1
	}
	pre, cur := 1, 1
	for i := 2; i < x; i++ {
		temp := cur
		cur += pre
		pre = temp
	}
	return cur
}

//FibonacciV2 记忆化搜索版本
//需要占用更多内存
var fib = []int{0,1,1}
func FibonacciV2(x int) int {
	for len(fib) <= x {
		fib = append(fib, -1)
	}
	if fib[x] != -1 {
		return fib[x]
	} else {
		fib[x] = FibonacciV2(x-1) + FibonacciV2(x-2)
		return fib[x]
	}
}
