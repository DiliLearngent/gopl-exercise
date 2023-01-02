// question
// zh
// 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。
// 只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
    "bytes"
	"golang.org/x/net/html"
)

var id int

func main() {
    // Crawl the web breadth-first,
    // starting from the command-line arguments.
    hosts := []string{}
    for _, u := range os.Args[1:] {
        u, err := url.Parse(u)
        if err != nil {
            continue
        }
        hosts = append(hosts, u.Host)
    }
    depth := flag.Int("depth", 2, "crawl depth")
    path := flag.String("path", "C:/code/gocode/gopl-exercise/src/p5/p5-13/www", "save path")
    flag.Parse()
    if *depth <= 0 {
        fmt.Println("The Crawl Depth Should > 0")
        os.Exit(1)
    }
    exist, err := PathExists(*path)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v", err)
        os.Exit(1)
    }
    if exist {
        err := os.RemoveAll(*path)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %v", err)
            os.Exit(1)
        }
    }
    err = os.Mkdir(*path, os.ModePerm)
    if err != nil {
        fmt.Println("create dir faile")
        os.Exit(1)
    }
    breadthFirst(crawl, os.Args[1:], hosts, *path, *depth)
}

func crawl(url string, path string) []string {
    fmt.Println(url)
    list, err := ExtractAndDown(url, path)
    if err != nil {
        log.Print(err)
    }
    return list
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string, path string) []string, worklist []string, hosts []string, path string, depth int) {
    seen := make(map[string]bool)
    var d int
    for len(worklist) > 0 && d < depth {
        items := worklist
        worklist = nil
        for _, item := range items {
            if !seen[item] && checkDomain(item, hosts) {
                seen[item] = true
                worklist = append(worklist, f(item, path)...)
            }
        }
        d++
    }
}


func ExtractAndDown(url string, path string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
    }
    bodyRes, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error : %v", err)
        os.Exit(1)
    }
    resbody := io.NopCloser(bytes.NewReader(bodyRes))
    filename := path + "/" + strconv.Itoa(id) + ".txt"
    f, err := os.Create(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "create file fail. err: %v", err)
    } else {
        // writer := bufio.NewWriter(f)
        // reader := bufio.NewReader(resp.Body)
        // for {
        //     value, err := reader.ReadString('\n')
        //     if err == io.EOF {
        //         break
        //     }
        //     writer.WriteString(value)
        //     writer.Flush()
        // }
        // _, err := io.Copy(f, resp.Body)
        // if err != nil {
        //     fmt.Fprintf(os.Stderr, "io copy fail. err: %v", err)
        // }
        writer := bufio.NewWriter(f)
        writer.WriteString(string(bodyRes))
        writer.Flush()
        id++
        f.Close()
    }
    doc, err := html.Parse(resbody)
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    resp.Body.Close()
    var links []string
    visitNode := func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, a := range n.Attr {
                if a.Key != "href" {
                    continue
                }
                link, err := resp.Request.URL.Parse(a.Val)
                if err != nil {
                    continue // ignore bad URLs
                }
                links = append(links, link.String())
            }
        }
    }
    forEachNode(doc, visitNode, nil)
    return links, nil
}

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



func checkDomain(u string, hosts []string) bool {
    s, err := url.Parse(u)
    if err != nil {
        return false
    }
    for _, host := range hosts {
        if s.Host == host {
            return true
        }
    }
    return false
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
