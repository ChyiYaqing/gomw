package main

import (
	"fmt"
	"log"
	"net/http"
)

// 实现ServerHTTP 接口实例

// Engine is the url handler for all requests
type Engine struct{}

// responsewriter 请求的响应; request HTTP请求的所有信息
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":1207", engine))
}
