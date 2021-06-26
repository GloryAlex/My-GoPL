package main

import (
	"GoPl/2.structure/wordfreq"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	freq := wordfreq.WordFreq(scanner)
	for s, i := range *freq {
		fmt.Printf("\"%s\" = %d\n", s, i)
	}
}
