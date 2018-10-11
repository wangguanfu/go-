package main

import (
	"fmt"
	"time"
	"os"
)

func mm() {
	for i:=0; i<5;i++{
		fmt.Println("11111111111111")
		time.Sleep(100 * time.Millisecond)
		if i == 2 {
			//return
			//runtime.Goexit()
			os.Exit(-24)
		}
	}
}

func test() {
	mm()
	//runtime.Goexit()  // 将 go 程 结束
	defer fmt.Println("cccc")
	fmt.Println("dddd")
}

func main()  {

	fmt.Println("aaaa")
	go test()
	time.Sleep(10 * time.Millisecond)

	fmt.Println("bbbb")

	for {
		fmt.Println("2222222222")
		time.Sleep(300 * time.Millisecond)
	}
}
