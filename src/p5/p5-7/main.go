// questions
// zh
// 完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。
// 使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。
// 编写测试，验证程序输出的格式正确。（详见11章）
package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
	"net/http"
	"strings"
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


var depth int
func startElement(n *html.Node) {
    if n.Type == html.ElementNode {
        fmt.Printf("%*s<%s", depth*2, "", n.Data)

		for _, attr := range n.Attr {
			fmt.Printf(" %s=%q", attr.Key, attr.Val)
		}
		if n.Data == "img" {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
		}
        depth++
    }else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}else if n.Type == html.TextNode && n.Parent.Data != "style" && n.Parent.Data != "script" {
		for _, line := range strings.Split(n.Data, "\n") {
			line = strings.TrimSpace(line)
			if line != "" && line != "\n" {
				fmt.Printf("%*s%s\n", depth*2, "", line)
			}
		}
	}
}

func endElement(n *html.Node) {
    if n.Type == html.ElementNode {
        depth--
		if n.Data != "img" || n.FirstChild != nil {
        	fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
    }
}