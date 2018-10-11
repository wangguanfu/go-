package main

import (
	"math/rand"
	"time"
	"fmt"
	"runtime"
	"sync"
)

// 创建条件变量
var cond sync.Cond

func producer2(out chan<- int, idx int)  {
	for {
		// 加条件变量中的锁
		cond.L.Lock()
		for len(out) == 5 {
		//if len(out) == 5 {
				cond.Wait()
		}
		// 访问共享数据区
		num := rand.Intn(500)
		out <- num
		fmt.Printf("-----%dth 生产者，生产：%d\n", idx,  num)
		// 解锁
		cond.L.Unlock()
		// 唤醒阻塞在条件变量上的对端 —— 消费者
		cond.Signal()
		//time.Sleep(time.Millisecond * 500)
	}
}
func consumer2(in <-chan int, idx int)  {
	for {
		cond.L.Lock()
		//if len(in) == 0 {
		for len(in) == 0 {
			cond.Wait()
		}
		num := <- in
		fmt.Printf("%dth 消费者，消费：%d\n", idx,  num)
		cond.L.Unlock()
		cond.Signal()

		//time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 5)

	// 初始化 cond
	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go producer2(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		go consumer2(ch, i+1)
	}

	for {
		runtime.GC()
	}
}


