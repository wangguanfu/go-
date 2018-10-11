package main

import (
	"time"
	"fmt"
)

func main()  {
	myTimer := time.NewTimer(time.Second * 5)

	go func() {
		<-myTimer.C
		fmt.Println("定时时间到！")
	}()
	//myTimer.Stop()
	myTimer.Reset(time.Second * 1)
	for {
		;
	}
}
