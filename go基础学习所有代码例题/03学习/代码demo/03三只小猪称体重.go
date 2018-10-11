package main

import (
	"fmt"
)

func main() {
	//有三只小猪 判断哪个最重

	var a, b, c int

	fmt.Scan(&a, &b, &c) //10  12  8

	if a > b {
		if a > c {
			fmt.Println("a最重")
		} else {
			fmt.Println("c最重")
		}
	} else {
		if b > c {
			fmt.Println("b最重")
		} else {
			fmt.Println("c最重")
		}
	}


}