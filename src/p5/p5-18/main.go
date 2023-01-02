// question
// zh
// 不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件。
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("%s content save to %s, byte %d\n", url, local, n)
	}
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", 0, err
    }
    defer resp.Body.Close()
    local := path.Base(resp.Request.URL.Path)
    if local == "/" {
        local = "index.html"
    }
    f, err := os.Create(local)
    if err != nil {
        return "", 0, err
    }
    //先调用io copy err
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
    n, err = io.Copy(f, resp.Body)
    return local, n, err
}
