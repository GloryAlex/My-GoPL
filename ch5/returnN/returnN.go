package main

import "fmt"

func main() {
	a := returnN()
	fmt.Print(a)
}

func returnN() (result int) {
	defer func() {
		switch p := recover(); p {
		case nil:
		default:
			result = p.(int)
		}
	}()
	panic(3)
}
