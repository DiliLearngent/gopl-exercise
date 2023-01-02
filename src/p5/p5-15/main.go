// question
// zh
// 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	//1
	if v, err := max(); err == nil {
		fmt.Printf("max: %d\n", v)
	} else {
		log.Println(err)
	}

	//2
	if v, err := min(1,2,3,4); err == nil {
		fmt.Printf("min: %d\n", v)
	} else {
		log.Println(err)
	}

	//3
	v := max1(4,5,6,7,8,9)
	fmt.Printf("max: %d\n", v)

	//4
	v = min1(4,5,6,7,8,9)
	fmt.Printf("min: %d\n", v)

}

func max(vals ...int) (int, error){
	if len(vals) == 0 {
		return 0, errors.New("no args")
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("no args")
	}
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min, nil
}

// at least recv 1 args
func max1(v int, vals ...int) int {
	max := v
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min1(v int, vals ...int) int {
	min := v
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}