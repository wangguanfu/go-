package main

import (
	"fmt"
	"time"
)

func main()  {

	ch := make(chan int, 3)
	fmt.Println("len(ch)=", len(ch), "cap(ch)=", cap(ch))

	// 创建匿名子go程
	go func() {
		fmt.Println("子 go 程 start ....")
		for i:=0; i<7;i++ {
			ch <- i
			fmt.Println("写---i = ", i, "len(ch)=", len(ch), "cap(ch)=", cap(ch))
		}
	}()

	time.Sleep(time.Second * 1)

	for i:=0; i<7; i++ {
		num := <-ch
		fmt.Println("----读：", num, "len(ch)=", len(ch), "cap(ch)=", cap(ch))
	}
}
