package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <-chan int, quit <-chan bool)  {
	for {
		select {
		case num := <-ch:			// 打印一个 fibonacci 数值
			fmt.Print(num, " ")
		case <-quit :				// 接收到go程结束通知
			//return
			runtime.Goexit()
		}
	}
}
func main()  {
	ch := make(chan int)
	quit := make(chan bool)
	// 创建一个打印fibonacci数列的 子go程
	go fibonacci(ch, quit)
	x, y := 1, 1
	for i:=0; i<50; i++ {
		ch <- x
		x, y = y, x+y
	}
	// 写完 20 个数，通知子go程结束
	quit <- false
}
