package main

import (
	"net/http"
	"fmt"
)

// 当浏览器访问服务器时，该函数自动被调用
 func handlerSRV(w http.ResponseWriter,  r *http.Request) {
 	w.Write([]byte("hello http"))

 	fmt.Println("Method:", r.Method)
 	fmt.Println("RemoteAddr:", r.RemoteAddr)
 	fmt.Println("Proto:", r.Proto)
 	fmt.Println("Header:", r.Header)
 	fmt.Println("URL:", r.URL)
 	fmt.Println("Body:", r.Body)
 }

func main()  {
	// 1. 注册处理函数
	http.HandleFunc("/", handlerSRV)

	// 2. 绑定服务器地址结构
	http.ListenAndServe("127.0.0.1:8000", nil)
}
