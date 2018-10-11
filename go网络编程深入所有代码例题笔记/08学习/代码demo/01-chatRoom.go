package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)
// 创建全局用户结构体类型
type Client struct {
	Name string
	Addr string
	C chan string
}
// 定义全局 map，存储在线用户列表
var OnlineMap map[string]Client

// 定义全局channel，完成用户广播消息
var message = make(chan string)

func WriteMsg2Client(clnt Client, conn net.Conn)  {
	// 监听用户de channel， 有数据读，没数据阻塞
	for msg := range clnt.C {
		// 读到消息，写给客户端。
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(clnt Client, msg string) (buf string) {
	 buf = "[ " + clnt.Addr + " ]" + clnt.Name + " : " + msg + "\n"
	 return
}

func HandlerConnect(conn net.Conn)  {
	defer conn.Close()
	// 获取连接上来的客户的 IP+port
	netAddr := conn.RemoteAddr().String()

	// 组织信息，初始化用户结构体, 初始时 Name == Addr == IP+port
	clnt := Client{netAddr, netAddr, make(chan string)}

	// 将新用户添加到 在线用户列表
	OnlineMap[netAddr] = clnt

	// 创建 go 程， 读取用户自带 channel, 获取用户广播消息。
	go WriteMsg2Client(clnt, conn)

	// 组织用户上线消息，写入全局 channel
	//message <- "[ " + netAddr + " ]" + clnt.Name + " : login"
	message <- MakeMsg(clnt, "login")

	// 创建 监听用户是否下线的 channel
	isQuit := make(chan bool)

	// 创建一个判断用户是否 有通信消息的 channel
	hasChat := make(chan bool)

	// 创建匿名go程，专门负责用户聊天信息的读取和  广播（将消息写到全局channel）
	go func() {
		buf := make([]byte, 4096)
		for {					// 循环监听 conn 上是否有客户端数据。
			n, err := conn.Read(buf)		// 读取客户端的数据，存入buf中
			if n == 0 {
				isQuit<-true 	// 用户下线
				fmt.Printf("检查到客户端 %s 关闭\n", clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err:", err)
				return
			}
			//fmt.Println("msg:", buf[:n])
			msg := string(buf[:n-1])	// 去除最后的 '\n' 字符

			// 判断用户发送的是一个 “查询在线用户命令”
			if msg == "who" && len(msg) == 3 {
				// 遍历在线用户列表 OnlineMap, 将提取到的用户信息，写给当前客户端
				for _, clnt := range OnlineMap {
					clntInfo := MakeMsg(clnt, "[OnLine]")
					conn.Write([]byte(clntInfo))
				}
			} else if len(msg) >= 8 && msg[:7] == "rename|" {		// 判断用户发送的是一个 “改名命令”
				// 提取用户自己命名的新名
				newName := strings.Split(msg, "|")[1]		// “rename”[0] “AAA”[1]
				// 将新用户名修改至在线用户列表。
				clnt.Name = newName				// 替换旧明
				OnlineMap[netAddr] = clnt		// 更新 client 信息。
				conn.Write([]byte("Rename successful !!!\n"))
			} else {
				// 将读到的数据内容写 到全局channel 中 —— 广播
				message <- MakeMsg(clnt, msg)
			}
			// 3个if分支有任意一个，说明用户活动。 重置超时计时。
			hasChat <- true
		}
	}()

	// 防止 当前 go 程结束
	for {
		select {
			case <-isQuit :				// 没有数据写入，阻塞等待。
				// 从 OnlineMap 摘除该用户
				delete(OnlineMap, netAddr)
				// 组织用户下线消息,  广播
				message <- MakeMsg(clnt, "logout")

				// 关闭 用户自带 channel —— 结束 WriteMsg2Client go程
				close(clnt.C)
				return
			case <-hasChat:			// 重置下面的时间
				// 啥也不做

			case <-time.After(time.Second * 10):
				// 从 OnlineMap 摘除该用户
				delete(OnlineMap, netAddr)
				// 组织用户下线消息,  广播
				message <- MakeMsg(clnt, "time out to Leave")
				// 关闭 用户自带 channel —— 结束 WriteMsg2Client go程
				close(clnt.C)
				return
		}
	}
}

func Manager()  {
	// 初始化全局在线用户列表
	OnlineMap = make(map[string]Client)

	// 循环 读全局channel
	for {
		// 无数据阻塞，有数据 遍历在线用户列表，写给每一个在线用户
		msg := <-message

		// 遍历在线用户列表
		for _, clnt := range OnlineMap {
			clnt.C <-msg		// 将从全局message中读到的数据，写给每一个用户的 channel
		}
	}
}

func main()  {
	// 创建监听socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Linten err:", err)
		return
	}
	defer listener.Close()

	// 创建管理者go程，管理全局channel的读事件 和 在线用户列表
	go Manager()

	// 循环监听客户端连接请求。
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}
		// 创建 go 程 管理客户端数据通信事件
		go HandlerConnect(conn)
	}
}
