// question
// zh
// gopl.io/ch5/outline2（5.5节）的startElement和endElement共用了全局变量depth，将它们修改为匿名函数，使其共享outline中的局部变量。
package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		err := outline(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "html tree :%v\n", err)
		}
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	var depth int
	var startElement func(n *html.Node)
	var endElement func(n *html.Node)
	startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	endElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

// forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    if post != nil {
        post(n)
    }
}


// var depth int
// func startElement(n *html.Node) {
//     if n.Type == html.ElementNode {
//         fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
//         depth++
//     }
// }

// func endElement(n *html.Node) {
//     if n.Type == html.ElementNode {
//         depth--
//         fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
//     }
// }