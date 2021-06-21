package intset

import (
	"fmt"
)

// IntSet
//位图数据结构
//第i位为0时表示数字i不存在
//否则表示数字i存在
type IntSet struct {
	words []int64
	size  int
}

func (s *IntSet) Len() int {
	return s.size
}

func (s *IntSet) Contains(x int) bool {
	arrayPosition, bitPosition := x/64, uint(x%64)
	return len(s.words) > arrayPosition && s.words[arrayPosition]&(1<<bitPosition) != 0
}

func (s *IntSet) Add(x int) {
	arrayPosition, bitPosition := x/64, uint(x%64)
	s.ensureCapacity(arrayPosition)
	s.words[arrayPosition] |= 1 << bitPosition

	s.size++
}

func (s *IntSet) Remove(x int) {
	arrayPosition, bitPosition := x/64, uint(x%64)
	s.ensureCapacity(arrayPosition)
	s.words[arrayPosition] &= (1 << bitPosition) - 1

	s.size--
}

func (s *IntSet) ensureCapacity(position int) {
	for len(s.words) <= position {
		s.words = append(s.words, 0)
	}
}

func (s *IntSet) UnionWith(another *IntSet) {
	for i, val := range another.words {
		if i < len(s.words) {
			s.words[i] |= val
		} else {
			s.words = append(s.words, val)
		}
	}

	s.size = 0
	for _, word := range s.words {
		s.size += popCount(word)
	}
}

func (s *IntSet) String() string {
	return fmt.Sprintf("%v", s.Elems())
}

func (s *IntSet) Clear() {
	s.size = 0
	for i := 0; i < len(s.words); i++ {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() IntSet {
	var ret IntSet
	length := len(s.words)
	for i := length - 1; i >= 0 && s.words[i] == 0; i++ {
	}
	for i := 0; i < length; i++ {
		ret.words = append(ret.words, s.words[i])
	}
	return ret
}

func (s *IntSet) Elems() []int64 {
	var ret []int64
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				ret = append(ret, int64(i*64+j))
			}
		}
	}
	return ret
}

func popCount(x int64) int {
	count := 0
	for x > 0 {
		count++
		x &= x - 1
	}
	return count
}
