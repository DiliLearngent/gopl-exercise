package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gettextcontent: %v\n", err)
		os.Exit(1)
	}

	arr := []string{}
	gettext(&arr, doc)

	for idx, v := range arr {
		fmt.Printf("Text %d:%-15s\n", idx, v)
	}


}

func gettext(arr *[]string, n *html.Node) {
	if n.Type == html.ElementNode &&
        (n.Data == "style" || n.Data == "script") {
        return
    }

	if n.Type == html.TextNode {
		*arr = append(*arr, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		gettext(arr, c)
	} 
}