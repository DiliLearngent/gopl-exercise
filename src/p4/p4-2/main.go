// question
// zh
// 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
package main

import (
	"fmt"
	"flag"
	"crypto/sha256"
	"crypto/sha512"
)


func main() {
	f := flag.String("flag", "sha256", "flag = sha256 | sha384 | sha512")
	flag.Parse()
	var s string
	fmt.Print("inupt string:")
	fmt.Scanf("%s\n", &s)

	switch *f {
	case "sha256":
		fmt.Printf("%s sha256: %x\n", s, sha256.Sum256([]byte(s)))
	case "sha384":
		fmt.Printf("%s sha384: %x\n", s, sha512.Sum384([]byte(s)))
	case "sha512":
		fmt.Printf("%s sha512: %x\n", s, sha512.Sum512([]byte(s)))
	}
}