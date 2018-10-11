package main

import (
	"math/rand"
	"time"
	"fmt"
	"runtime"
)

func producer(out chan<- int, idx int)  {
	for {
		num := rand.Intn(500)
		out <- num
		fmt.Printf("-----%dth 生产者，生产：%d\n", idx,  num)
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(in <-chan int, idx int)  {
	for {
		num := <- in
		fmt.Printf("%dth 消费者，消费：%d\n", idx,  num)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go producer(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		go consumer(ch, i+1)
	}

	for {
		runtime.GC()
	}
}


