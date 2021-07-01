package main

import (
	"GoPl/ch8/rocket"
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort")

	ticker := time.NewTicker(1 * time.Second)
	for i := 10; i >= 0; i-- {
		select {
		case <-ticker.C:
			fmt.Println(i)
		case <-abort:
			fmt.Println("Launch abort!")
			ticker.Stop()
			return
		}
	}
	rocket.Launch()
}
