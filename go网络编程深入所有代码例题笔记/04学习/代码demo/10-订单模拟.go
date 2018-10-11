package main

import "fmt"

type orderInfo struct {
	id int
}
// 生产者
func makeOrder(make chan<- orderInfo)  {
	for i:=0; i<5; i++ {
		order := orderInfo{i+1}			// 模拟产生订单
		make <- order		//将订单放入单的缓冲区
	}
	close(make)
}
// 消费
func dealOrder(deal <-chan orderInfo) {
	for order := range deal {
		fmt.Println("处理掉订单：", order.id)
	}
}

func main()  {

	ch := make(chan orderInfo, 4)		// 存放订单的缓冲区

	go makeOrder(ch)

	dealOrder(ch)
}
