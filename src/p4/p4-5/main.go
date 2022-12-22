// question
// zh
// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
package main 

import (
	"fmt"
)

func main() {
	s := []string{"aaa", "bbb", "bbb", "ccc", "aaa"}
	r := remove(s)
	fmt.Printf("%q\n", r)
}

func remove(s []string) []string {
	for i, v := range s {
		for j := i+1; j < len(s); j++ {
			if v == s[j] {
				copy(s[j:], s[j+1:])
				s = s[:len(s)-1]
			}
		}
	}
	return s
}


// 消除相邻重复字符串
// func remove(s []string) []string {
// 	for i := 0; i < len(s)-1; {
// 		if s[i] == s[i+1] {
// 			copy(s[i:], s[i+1:])
// 			s = s[:len(s)-1]
// 		} else {
// 			i++
// 		}
// 	}
// 	return s
// }
