package main

import (
	"net/http"
	"os"
	"fmt"
)

func OpenAndSendFile(fileName string, w http.ResponseWriter)  {
	// 只读方式打开文件。
	path := "C:/itcast/test"		// 等价于："C:\\itcast\\test"
	f, err := os.Open(path+fileName)
	if err != nil {
		fmt.Println("Open err:", err)
		//w.Write([]byte("no such file or directroy!"))
		w.Write([]byte("<html><head><title>404 Not Found</title></head>" +
			"<body bgcolor=\"#cc99cc\"><h4>404 Not Found</h4>NO SUCH FILE OR DIR IS HERE!!!<hr>" +
			"</body></html>"))
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			fmt.Println("读取文件完毕\n")
			break
		}
		w.Write(buf[:n])		// 借助 ResponseWriter 将读到的数据内容，写回给 浏览器
	}
}
func handConn( w http.ResponseWriter, r *http.Request)  {
	// 获取用户请求的文件名
	fileName := r.URL.String()
	// 创建函数尝试读取文件内容，回发给浏览器
	OpenAndSendFile(fileName, w)
}

func main()  {
	http.HandleFunc("/", handConn)
	http.ListenAndServe("127.0.0.1:8006", nil)
}
