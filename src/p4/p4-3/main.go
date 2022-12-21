// question
// zh
// 重写reverse函数，使用数组指针代替slice。
package main

import (
	"fmt"
)

func main() {
	arr := [7]int{4, 6, 8, 9, 34, 23, 90}
	reverse(&arr)
	fmt.Printf("reverse array: %v\n", arr)
}

func reverse(arr *[7]int) {
	for i, j := 0, len(arr) - 1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}