package main

import (
	"fmt"
	"time"
)
func singing()  {
	for i:= 0; i<5; i++ {
		fmt.Println("======正在唱：青藏高原=====")
		time.Sleep(time.Second)
	}
}

func dancing()  {
	for i:= 0; i<5; i++ {
		fmt.Println("======正在跳：赵四街舞=====")
		time.Sleep(time.Second)
	}
}

func main()  {
	go singing() // 子 go 程  --  Goroutine  --- Coroutine
	go dancing()
/*	for {			// 作用是，保证主go程不结束 —— 进程地址空间 不会消失。
		;
	}*/
}
