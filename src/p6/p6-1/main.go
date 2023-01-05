// question
// zh
// 为bit数组实现下面这些方法
// func (*IntSet) Len() int      // return the number of elements
// func (*IntSet) Remove(x int)  // remove x from the set
// func (*IntSet) Clear()        // remove all elements from the set
// func (*IntSet) Copy() *IntSet // return a copy of the set
package main

import (
	"fmt"
	"bytes"
)
type IntSet struct {
	words []uint64
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println("x:",x.String())
	fmt.Println("IntSet x len:", x.Len())
	y := x.Copy()
	fmt.Println("y:", y.String())
	x.Remove(144)
	fmt.Println("after remove x:", x.String())
	x.Clear()
	fmt.Println("after clear x:", x.String())
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
    word, bit := x/64, uint(x%64)
    return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
    word, bit := x/64, uint(x%64)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < 64; j++ {
            if word&(1<<uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", 64*i+j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}

// return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for bit := uint(0); bit < 64; bit ++ {
			if word & (1<<bit) != 0 {
				count++
			}
		}
	}
	return count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/64, uint(x%64)
		s.words[word] &^= 1 << bit
	}
}

// remove all elements from the set
func (s * IntSet) Clear() {
	s.words = nil
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var set IntSet
	set.words = append(set.words, s.words...)
	return &set
}


