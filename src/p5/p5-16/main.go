// question
// zh
// 编写多参数版本的strings.Join。
package main

import (
	"fmt"
)

func main() {
	s := Join("|", "aaa", "bbb", "ccc")
	fmt.Printf("%s\n", s)
	s = Join("|", "aaa")
	fmt.Printf("%s\n", s)
	s = Join("|")
	fmt.Printf("%s\n", s)
}

func Join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else {
		str := ""
		for i := 0; i < len(strs) -1; i++ {
			str += strs[i]
			str += sep
		}
		str += strs[len(strs)-1]
		return str
	}
}