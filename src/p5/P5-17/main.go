// question
// zh
// 编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Println(err)
		}
		resp.Body.Close()
		images := ElementsByTagName(doc, "img")
		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
		for _, image := range images {
			fmt.Print("<img")
			for _, attr := range image.Attr {
				fmt.Printf(" %s=%q", attr.Key, attr.Val)
			} 
			fmt.Println("></img>")
		}
		fmt.Println("######################")
		for _, heading := range headings {
			fmt.Println(heading.Data)
			fmt.Println(heading.Attr)
			fmt.Println("----------------------------------")
		}
	}
}

func ElementsByTagName(doc *html.Node, name...string) []*html.Node {
	var nodes []*html.Node
	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, na := range name {
				if n.Data == na {
					nodes = append(nodes, n)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)
	return nodes
}