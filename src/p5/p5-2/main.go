package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"

)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getelementcount: %v\n", err)
		os.Exit(1)
	}

	m := map[string]int{}
	getcount(m, doc)

	for k, v := range m {
		fmt.Printf("%-s's count:%-5d\n", k, v)
	}
}

func getcount(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getcount(m, c)
	}
}