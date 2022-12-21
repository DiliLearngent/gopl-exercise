// question
// zh
// 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func main() {
	str := os.Args[1:]
	if len(str) != 2 {
		fmt.Println("input two args")
		os.Exit(1)
	}

	c1 := sha256.Sum256([]byte(str[0]))
	c2 := sha256.Sum256([]byte(str[1]))

	count := diffbitcount(c1, c2)
	fmt.Printf("%s and %s sha256 different bit count:%d\n", str[0], str[1], count)
}

func diffbitcount(s1, s2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		temp := s1[i] ^ s2[i]
		count += int(pc[temp])
	}
	return count
}