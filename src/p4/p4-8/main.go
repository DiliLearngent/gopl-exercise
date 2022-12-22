// question
// zh
// 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。
package main

import (
	"fmt"
	"unicode"
	"os"
	"bufio"
	"io"
)

func main() {
	counts := make(map[string]int, 5)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
		}
		switch{
		case unicode.IsLetter(r):
			counts["Letter"]++
		case unicode.IsNumber(r):
			counts["Number"]++
		case unicode.IsGraphic(r):
			counts["Graphic"]++
		case unicode.IsSymbol(r):
			counts["Symbol"]++
		case unicode.IsSpace(r):
			counts["Space"]++
		}
	}
	for class, count := range counts {
		fmt.Printf("%s count: %d\n", class, count)
	}
} 