// question
// zh
// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
package main 

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	b := []byte("sdgsdg  sd哈哈  哈g")
	r := removespace(b)
	fmt.Printf("%q\n", r)
}

func removespace(b []byte) []byte {
	for i := 0; i < len(b); {
		value1, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(value1) {
			if value2, _ := utf8.DecodeRune(b[i+size:]); unicode.IsSpace(value2) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
			}
		}
		i += size
	}
	// for i := 0; i < len(b)-1; i++ {
	// 	if unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[i+1])){
	// 		copy(b[i:], b[i+1:])
	// 		b = b[:len(b)-1]
	// 	}
	// }
	return b
}