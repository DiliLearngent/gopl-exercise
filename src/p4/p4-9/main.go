// question
// zh
// 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := map[string]int{}
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%s count: %d\n", k, v)
	}
}