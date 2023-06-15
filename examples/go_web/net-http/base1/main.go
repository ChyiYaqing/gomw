package main

import (
	"fmt"
	"log"
	"net/http" // 提供基本的Web功能，监听端口，映射静态路由，解析HTTP报文。
)

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:1207", nil))
}

/*
~ via 🐹 v1.20.4 took 6m 25s
➜ curl http://localhost:1207/
URL.Path = "/"

~ via 🐹 v1.20.4
➜ curl http://localhost:1207/hello
Header["User-Agent"] = ["curl/8.1.2"]
Header["Accept"] = ["*/*"]
*/
