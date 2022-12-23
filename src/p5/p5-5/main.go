// question
// zh
// 实现countWordsAndImages。（参考练习4.9如何分词）
package main

import (
	"fmt"
	"net/http"
	"os"
    "strings"
	"golang.org/x/net/html"
)

func main() {
    urls := os.Args[1:]
    for _, url := range urls {
        words, images, err := CountWordsAndImages(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("%s:  (words:%d, images:%d)\n", url, words, images)
    }
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}
func countWordsAndImages(n *html.Node) (words, images int) {
    if n.Type == html.ElementNode && n.Data == "img" {
        images++
    }
    if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
        return
    }
    if n.Type == html.TextNode {
        text := strings.TrimSpace(n.Data)
        for _, l := range strings.Split(text, "\n") {
            words += len(strings.Split(l, " "))
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        w, i := countWordsAndImages(c)
        words += w
        images += i
    }
    return
}