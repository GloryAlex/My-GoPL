package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var wordCounts = make(map[string]int)
	var input = bufio.NewScanner(os.Stdin)
	for input.Scan() {
		wordCounts[input.Text()]++
	}
	for word, frequency := range wordCounts {
		if frequency > 1 {
			fmt.Printf("%d\t%s\n", frequency, word)
		}
	}
}
