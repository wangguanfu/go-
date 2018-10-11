package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)		// 用来进行数据通信
	quit := make(chan bool)		// 用来判断是否结束go程的channel

	go func() {
		for i:=0; i<5; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 300)
		}
		close(ch)
		fmt.Println("子go程打印完毕")
		quit <- true			// 通知主go程可以终止
	}()
	for {
		select {
		case num := <-ch :
			fmt.Println("num = ", num)
		case <-quit :
			//time.Sleep(time.Second)
			fmt.Println("主go程 收到子消息， 结束运行")
			//break		// 只能跳出当前 的 select
			return
			//runtime.Goexit()	// 退出主 go 程 ———— 不会清理进程地址空间。
			//os.Exit(0)
		}
		fmt.Println("==========================")
	}
}
