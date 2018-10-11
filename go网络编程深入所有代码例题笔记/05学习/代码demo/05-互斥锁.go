package main

import (
	"fmt"
	"time"
	"sync"
)
//var ch = make(chan int)
var mutex sync.Mutex		// 创建一个互斥锁（互斥量）

func printer(s string)  {
	mutex.Lock()				// 访问共享数据前 加锁。
	defer mutex.Unlock()		// 访问共享数据结束，解锁。
	for _, char := range s {
		fmt.Printf("%c", char)				// stdout
		time.Sleep(time.Millisecond * 300)
	}
	fmt.Println("")
}

func person1()  {
	printer("hello")
}

func person2()  {
	printer("world")
}

func main()  {
	go person1()
	go person2()

	for {
		;
	}
}
