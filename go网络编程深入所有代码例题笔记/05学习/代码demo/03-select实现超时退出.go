package main

import (
	"fmt"
	"time"
)

func main()  {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {

		for {
			select {
				case num:=<-ch:
					fmt.Println("num = ", num)
				case <-time.After(time.Second * 3):
					fmt.Println("time out")
					quit <- true
					//return		// Goexit(）  等价
					goto ABC		// lable:标签名， 改名字任意
			}
		}
	ABC:		// 不能超出当前函数。
		fmt.Println("----break to lable----")
	}()

	for i:=0; i<2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

	<-quit		// 阻塞
	fmt.Println("主 go 程 执行结束。")
}
