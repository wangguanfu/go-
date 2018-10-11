package main

import (
	"net/http"
	"fmt"
	"io"
)

/*func main()  {
	// 获取服务器数据，
	resp, err := http.Get("http://127.0.0.1:8006/lf.png")
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()  // resp.close ---- bool

	fmt.Println("Status:", resp.Status)
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("Header:", resp.Header)
	fmt.Println("Proto:", resp.Proto)
	fmt.Println("Body:", resp.Body)

}*/

func main()  {
	// 获取服务器数据，
	resp, err := http.Get("http://127.0.0.1:8006/lf.txt")
	//resp, err := http.Get("https://www.sina.com.cn/")
	//resp, err := http.Get("https://www.baidu.com/")
	//resp, err := http.Get("http://www.itcast.cn/")
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()  // resp.close ---- bool

	buf := make([]byte, 4096)

	var result string

	//循环从 Body 中 读取数据内容。
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err == io.EOF {
			fmt.Println("err:", err)
		}
		result += string(buf[:n])
	}
	fmt.Println("resutl:", result)
}