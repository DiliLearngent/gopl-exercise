// question
// zh
// 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环。
package main

import (
	"fmt"
	"os"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
    "algorithms": {"data structures"},
    "calculus": {"linear algebra"},
    "compilers": {
        "data structures",
        "formal languages",
        "computer organization",
    },
    "data structures":       {"discrete math"},
    "databases":             {"data structures"},
    "discrete math":         {"intro to programming"},
    "formal languages":      {"discrete math"},
    "networks":              {"operating systems"},
    "operating systems":     {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
	"linear algebra": {"calculus"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	for i, v := range order {
		fmt.Printf("%d:%s\n", i+1, v)
	}
}

func topoSort(m map[string][]string) ([]string, error){
	order := []string{}
	seen := make(map[string]bool)
	stack := make(map[string]bool)
	var visitAll func(items []string) error
	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				stack[item]= true
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				stack[item] = false
				order = append(order, item)
			} else if stack[item]{
					return fmt.Errorf("cycle:%s", item)
			}
		}
		return nil
	}

	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, err
	}
	return order, nil
}