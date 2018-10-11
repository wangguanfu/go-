package main

import "net/http"

// 借助go 封装好的 访问http的标准库，创建 服务器。

func handler(w http.ResponseWriter, r *http.Request)  {
	// w: 表示 写回给客户端（浏览器）的数据信息
	// r：表示从 浏览器读到的数据信息
	w.Write([]byte("hello http 123"))
}

func main()  {
	// 1. 注册处理函数（回调函数）
	http.HandleFunc("/itcast", handler)

	// 2. 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8007", nil)
}
