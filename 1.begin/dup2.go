package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		return
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v", err) //ignore error
			}
			wordCounts := countLines(f)
			if len(wordCounts) > 0 {
				fmt.Printf("Duplicates in \"%v\":\n", file)
				for word, freq := range wordCounts {
					fmt.Printf("%d\t%s\n", freq, word)
				}
			}
		}
	}
}

func countLines(f *os.File) map[string]int {
	wordCounts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		wordCounts[input.Text()]++
	}
	for word, freq := range wordCounts {
		if freq <= 1 {
			delete(wordCounts, word)
		}
	}
	return wordCounts
}
