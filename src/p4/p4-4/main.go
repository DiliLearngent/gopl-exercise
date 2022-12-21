// question
// zh
// 编写一个rotate函数，通过一次循环完成旋转。
package main

import (
	"fmt"
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	r := rotate(s, 2)
	fmt.Printf("rotate result: %v\n", r)
}

func rotate(s []int, n int) []int {
	r := []int{}
	for i := n-1; i >= 0; i-- {
		r = append(r, s[i])
	}
	r = append(r, s[n:]...)
	return r
}