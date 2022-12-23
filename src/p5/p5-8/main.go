// question
// zh
// 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。
// 使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。
// func ElementByID(doc *html.Node, id string) *html.Node
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)
func main() {
	var s string
	for _, url := range os.Args[1:] {
		fmt.Print("input node id to find:")
		fmt.Scanf("%s\n", &s)
		err := outline(url, s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "find node by id: %v", err)
		}
	}
}

func outline(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	if n := ElementByID(doc, id); n != nil {
		fmt.Println(n.Data)
		for _, attr := range n.Attr {
			fmt.Printf("%s=%q\n", attr.Key, attr.Val)
		}
	}
	return nil

} 

func ElementByID(doc *html.Node, id string) *html.Node {
	if node := forEachNoded(doc, id, foundId, nil); node != nil {
		return node
	} else {
		return nil
	}
}

func forEachNoded(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) (*html.Node){
	if pre != nil && pre(n, id) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node := forEachNoded(c, id, pre, post); node != nil {
			return node
		}
	}
	if post != nil && post(n, id) {
		return n
	}
	return nil
}

func foundId(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
	}
	return false
}