package wordfreq

import "bufio"

func WordFreq(scanner *bufio.Scanner) *map[string]int {
	dict := make(map[string]int)
	for scanner.Scan() {
		dict[scanner.Text()]++
	}
	return &dict
}
