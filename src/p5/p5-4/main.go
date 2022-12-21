// question
// zh
// 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink2: %v\n", err)
		os.Exit(1)
	}

	// the target can also global
	target := map[string]string{"a": "href", "img": "src", "script": "src", "link": "href"}
	links := map[string][]string{"a": {}, "img": {}, "script": {}, "link": {}}

	for k, v := range visit(target, links, doc) {
		fmt.Printf("%s's link:\n", k)
		for _, vv := range v {
			fmt.Printf("%s\n", vv)
		}
	}
}

func visit(target map[string]string, links map[string][]string, n *html.Node) map[string][]string {
	if n.Type == html.ElementNode {
		if v, ok := target[n.Data]; ok {
			for _, a := range n.Attr {
				if a.Key == v {
					links[n.Data] = append(links[n.Data], a.Val)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(target, links, c)
	}
	return links
}
