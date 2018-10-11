package main

import (
	"fmt"
	"runtime"
)

func main()  {

	go func(s string) {
		for i:=0; i<2;i++ {
			fmt.Println(s)
		}
	}("hello")

	for i:=0; i<2; i++ {
		runtime.Gosched()		// 主动让出 cpu 使用权
		fmt.Println("world")
	}
}
