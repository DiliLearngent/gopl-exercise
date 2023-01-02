// question
// zh
// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
package main

import (
	"fmt"
	"log"
)

func main() {
	result := f()
	fmt.Println(result)
}

func f() (result int) {
	defer func ()  {
		if p := recover(); p != nil {
			log.Println("Recover f", p)
		}
		result = 3
	}()
	panic("error")
}