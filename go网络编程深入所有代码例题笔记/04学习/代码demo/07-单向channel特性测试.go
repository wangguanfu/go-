package main

import "fmt"

func main()  {
	ch :=make(chan int)
	var send chan <- int		// 单向写channel
	send = ch
	send <- 907
	//<- send

	var recv <-chan int = ch		// 单向读channel
	fmt.Println("recv:", <-recv)
	//recv <- 8987

	//var ch2 chan int = send
}
