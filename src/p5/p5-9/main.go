// question
// zh
// 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
// func expand(s string, f func(string) string) string
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello, foo. foo, how are you?"
	s = expand(s, replace)
	fmt.Println(s)
}

func expand( s string, f func (string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}

func replace(s string) string {
	return "Mike"
}