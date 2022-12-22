// question
// zh
// 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("abcd哈喽efgh")
	revreseutf8(b)
	fmt.Printf("%s\n", b)
}

func revreseutf8(b []byte) {
	for i := 0; i < len(b); {
		v, size := utf8.DecodeRune(b[0:len(b)-i])
		copy(b[0:len(b)-i-size], b[size:len(b)-i])
		utf8.EncodeRune(b[len(b)-i-size:len(b)-i], v)
		i += size
	}
}

// func main() {
// 	b := []byte("abcd哈喽efgh")
// 	reverseUTF8(b)
// 	fmt.Printf("%s\n", b)
// }

// func reverseUTF8(b []byte) {
// 	for i := 0; i < len(b); {
// 		_, size := utf8.DecodeRune(b[i:])
// 		reverse(b[i : i+size])
// 		i += size
// 	}
// 	reverse(b)
// }

// func reverse(b []byte) {
// 	last := len(b) - 1
// 	for i := 0; i < len(b)/2; i++ {
// 		b[i], b[last-i] = b[last-i], b[i]
// 	}
// }