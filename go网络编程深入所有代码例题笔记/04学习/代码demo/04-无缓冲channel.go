package main

import (
	"fmt"
	"time"
)

func main()  {

	// 在主、子go程之间进行数据传递。 指定了 读写数据的先后顺序。 子先，主后
	ch := make(chan int, 0)
	fmt.Println("len(ch)=", len(ch), "cap(ch)=", cap(ch))

	// 创建匿名子go程
	go func() {
		fmt.Println("子 go 程 start ....")
		for i:=0; i<5;i++ {
			ch <- i
			fmt.Println("子go程 i = ", i, "len(ch)=", len(ch), "cap(ch)=", cap(ch))
			time.Sleep(time.Second)
		}
	}()

	for i:=0; i<5; i++ {
		num := <-ch
		fmt.Println("主go程读到：", num)
	}
}
