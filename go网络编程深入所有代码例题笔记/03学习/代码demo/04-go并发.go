package main

import (
	"fmt"
	"time"
)

func task()  {
	for {
		fmt.Println(" I am task")
		time.Sleep(300 * time.Millisecond)
	}
}

func main()  {

	// 创建子go程
	go task()

	go func () {
		for {
			fmt.Println("I am 匿名 goroutine")
			time.Sleep(300 * time.Millisecond)
		}
	}()

	i := 0
	// 主 go程的 执行指令
	for {
		i++
		fmt.Println("I am main goroutine, i=", i)
		time.Sleep(300 * time.Millisecond)
		if i == 3 {
			break
		}
	}
}
