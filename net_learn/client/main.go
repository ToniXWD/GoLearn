package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 使用 http.Get 发送一个 HTTP GET 请求。
	res, err := http.Get("http://127.0.0.1:8080/goodbye")
	if err != nil {
		panic(err)
	}

	// 使用 ioutil.ReadAll 读取响应体的全部内容，并关闭响应体。
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	// 打印响应的内容。
	fmt.Printf("%s", body)
}
