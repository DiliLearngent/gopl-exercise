// question
// zh
// (*IntSet).UnionWith会用|操作符计算两个集合的并集，
// 我们再为IntSet实现另外的几个函数IntersectWith（交集：元素在A集合B集合均出现），
// DifferenceWith（差集：元素出现在A集合，未出现在B集合），
// SymmetricDifference（并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A）。
package main

import (
	"fmt"
	"bytes"
)

type IntSet struct {
	words []uint64
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(4)
	x.Add(5)
	y.Add(3)
	y.Add(4)
	y.Add(6)
	//并集
	n := x.Copy()
	n.UnionWith(&y)
	fmt.Println("x union y:", n.String())
	//交集
	n = x.Copy()
	n.IntersectWith(&y)
	fmt.Println("x intersect y:", n.String())
	//差集
	n = x.Copy()
	n.DifferenceWith(&y)
	fmt.Println("x difference y:", n.String())
	//并差集
	n = x.Copy()
	n.SymmetricDifference(&y)
	fmt.Println("x symmetricdifference y:", n.String())
}

// 并集
// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

// 交集
func (s *IntSet) IntersectWith(t *IntSet) {
	l := len(s.words)
	if len(s.words) > len(t.words) {
		l = len(t.words)
	}
	for i := 0; i < l; i++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[:l]
}

//差集
func (s *IntSet) DifferenceWith(t *IntSet) {
	 for i, word := range s.words {
		if i >= len(t.words) {
			continue
		}
		s.words[i] = word &^ t.words[i]
	 }
}

// 并差集
func (s *IntSet) SymmetricDifference(t *IntSet) {
	u := s.Copy()
	s.UnionWith(t)
	u.IntersectWith(t)
	s.DifferenceWith(u)
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var set IntSet
	set.words = append(set.words, s.words...)
	return &set
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

