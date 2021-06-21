package main

import (
	"GoPl/ch6/intset"
	"fmt"
)

func main() {
	var x = intset.IntSet{}
	var y = intset.IntSet{}
	x.Add(1)
	x.Add(144)
	x.Add(45)
	y=x.Copy()
	x.Clear()
	fmt.Println(&x,&y)
}
