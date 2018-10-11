package main

import "fmt"

/*func main()  {
	ch := make(chan bool)
	go func() {
		for i:=0; i<3; i++ {
			ch <- true
		}
		close(ch)
		//ch <- 789
	}()
	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("读到 子 go 程数据：", num)
		} else {
			num := <-ch
			// 说明写端已经关闭了。
			fmt.Println("检测到 子go程关闭channel, num:", num)
			break
		}
	}
}*/

// 简单判别 channel 关闭方法
func main()  {
	ch := make(chan int)
	go func() {
		for i:=0; i<7; i++ {
			ch <- i
		}
		close(ch)
		//ch <- 789
	}()
/*	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("读到 子 go 程数据：", num)
		} else {
			// 说明写端已经关闭了。
			fmt.Println("检测到 子go程关闭channel, num:", num)
			break
		}
	}*/
	for num := range ch {
		fmt.Println("读到 子 go 程数据：", num)
	}
	fmt.Println("子go程数据读取完毕。channel关闭")
}