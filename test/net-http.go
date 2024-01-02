package main

import (
	"fmt"
	"html"
	"net/http"
)

//// 创建一个foo路由和处理函数
//http.Handle("/foo", fooHandler)
//
//// 创建一个bar路由和处理函数
//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//})
//
////监听8080端口
//log.Fatal(http.listenAndServe(":8080", nil))

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
}
