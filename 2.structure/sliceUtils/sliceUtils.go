package sliceUtils

import "unicode"

func Reverse(s []int) []int {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

func Rotate(s []int, k int) []int {
	k %= len(s)
	Reverse(s[:k])
	Reverse(s[k:])
	Reverse(s)
	return s
}

func RemoveDuplicateSpace(bytes []byte) []byte {
	ret := bytes[:0]
	for i := 0; i < len(bytes); i++ {
		if len(ret) > 0 && unicode.IsSpace(rune(ret[len(ret)-1])) && unicode.IsSpace(rune(bytes[i])) {
			continue
		}
		ret = append(ret, bytes[i])
	}
	return ret
}
