package main

// 死锁1
/*func main()  {
	ch := make(chan int)
	ch <- 88
	num := <-ch
	fmt.Println("num = ", num)
}*/

// 死锁2
/*func main()  {
	ch := make(chan int)
	go func() {
		num := <-ch	// 读
		fmt.Println("num = ", num)
	}()
	ch <- 99		// 写 —— 阻塞
	for {
		;
	}
}*/

// 死锁3
func main()  {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case <-ch1:
				ch2 <- 777
			}
		}
	}()
	for {
		select {
		case <-ch2 :
			ch1 <- 999
		}
	}
}
