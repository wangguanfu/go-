package main

import "fmt"

// 限定当前函数内部，只能做写操作。
func send(out chan<- int)  {
	out <- 666
	fmt.Println("send 666 over")
}

func recv(in <-chan int)  {
	num := <-in
	fmt.Println("recv num=", num)
}

func main()  {
	// 创建一个 双向 channel
	ch := make(chan int)

	go send(ch)
	go recv(ch)

	for {
		;
	}
}