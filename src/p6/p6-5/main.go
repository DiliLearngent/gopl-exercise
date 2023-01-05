// question
// zh
// 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
// 修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。
// 当然了，这里我们可以不用简单粗暴地除64，可以定义一个常量来决定是用32还是64，
// 这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
package main

import (
	"fmt"
	"bytes"
)

type IntSet struct {
	words []uint
}

const uintSize = 32 << (^uint(0) >> 63)

func main() {
	var x IntSet
	x.Add(56)
	fmt.Println("x:", x.String())
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
    word, bit := x/uintSize, uint(x%uintSize)
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
        for j := 0; j < uintSize; j++ {
            if word&(1<<uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", uintSize*i+j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}
