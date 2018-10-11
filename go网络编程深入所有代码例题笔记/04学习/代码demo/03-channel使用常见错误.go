package main

import (
	"fmt"
)

func main()  {

	cha := make(chan int)
	num := <- cha
	fmt.Println("读 到 channel 中的数据：", num)

	go func() {					// 子 go程
		cha <- 456
		fmt.Println("写 channel 完成")

		//time.Sleep(time.Second)
	}()
}
