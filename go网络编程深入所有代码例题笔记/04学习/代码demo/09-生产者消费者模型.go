package main

import "fmt"

func producer(send chan<- int)  {
	for i:=0; i<10; i++ {
		send <- i * i
		//fmt.Println("生产者 产生：", i*i)
	}
	close(send)
}
func consumer(recv <-chan int)  {
	for {
		if num, ok := <-recv; ok {
			fmt.Println("消费者 消费：", num)
		} else {
			break
		}
	}
/*	// 判断 channel 是否关闭。
	for num := range recv {
		fmt.Println("消费者 消费：", num)
	}*/
}
func main()  {
	// 创建缓冲区，带缓冲 双向 channel
	ch := make(chan int, 5)
	go producer(ch)		// 子go程 --- 生产者
	consumer(ch)		// 主 go 程 --- 消费者
}
