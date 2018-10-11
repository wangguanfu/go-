package main

import (
	"fmt"
	"time"
)

/*// 定义一个用于通信的 channel, 完成go程同步
var channel = make(chan int)

func printer(s string)  {
	for _, ch := range s {
		fmt.Printf("%c", ch)	 // 屏幕（硬件）—— 标准输出（文件）stdout
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("")
}
func person1()  {
	printer("hello")			// 先
	// 写 channel
	channel <- 666
}
func person2()  {
	// 读 channel
	<- channel
	printer("world")			// 后
}
func main()  {
	go person1()
	go person2()

	for {		// 防止主go程提早结束
		;
	}
}*/

// 定义一个用于通信的 channel, 完成go程同步
var channel = make(chan bool)

func printer(s string)  {
	for _, ch := range s {
		fmt.Printf("%c", ch)	 // 屏幕（硬件）—— 标准输出（文件）stdout
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("")
}
func person1()  {
	// 读 channel
	myBool := <- channel
	fmt.Println("person1 读到：", myBool)
	printer("hello")	// 后
}
func person2()  {
	printer("world")	// 先
	// 写 channel
	channel <- true
}
func main()  {
	go person1()
	go person2()

	for {		// 防止主go程提早结束
		;
	}
}
