package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type message struct {
	Msg string `json:"msg"`
}

// 自定义路由器
type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 前缀匹配

	if strings.HasPrefix(r.URL.Path, "/hello") {
		helloHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/goodbye") {
		goodbyeHandler(w, r)
		return
	}

	// 如果没有匹配的路径，返回404错误
	http.NotFound(w, r)
}

// 定义 helloHandler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	msg := message{Msg: "Hello, you've hit " + r.URL.Path}
	json.NewEncoder(w).Encode(msg)
}

// 定义 goodbyeHandler
func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	msg := message{Msg: "Hello, you've hit " + r.URL.Path}
	json.NewEncoder(w).Encode(msg)
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8080", mux)
}
