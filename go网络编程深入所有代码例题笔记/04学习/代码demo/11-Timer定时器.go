package main

import (
	"time"
	"fmt"
)

/*func main()  {
	// 打印当前时间
	fmt.Println("now:   ", time.Now())

	// 创建 Timer 定时器
	myTimer := time.NewTimer(time.Second * 2)

	time := <-myTimer.C		// 默认读——阻塞，定时时长到达后，系统写入当前时间，解除阻塞

	fmt.Println("定时后:", time)
}*/

func main()  {
	fmt.Println("now:   ", time.Now())

	time := <-time.After(time.Second * 2)

	fmt.Println("定时后:", time)
}