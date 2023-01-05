// question
// zh
// 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。
package main

import (
	"fmt"
)
type IntSet struct {
	words []uint64
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(1044)
	elems := x.Elems()
	fmt.Println(elems)
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
    word, bit := x/64, uint(x%64)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

// return a slice contain element of the set
func (s *IntSet) Elems() ([]int) {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for bit := 0; bit < 64; bit++ {
			if word & (1<<uint(bit)) != 0 {
				elems = append(elems, i * 64 + bit)
			}
		}
	}
	return elems
}