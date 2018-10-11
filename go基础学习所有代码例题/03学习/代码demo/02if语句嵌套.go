package main

import "fmt"

func main() {
	var score int

	fmt.Scan(&score)
	//if 语句嵌套
	if score > 700 {
		fmt.Println("我要上清华")

		if score > 720 {
			fmt.Println("我要学习计算机")
		} else if score > 710 {
			fmt.Println("我要学习挖掘机")
		} else {
			fmt.Println("我要学习美容美发")
		}

	} else if score > 680 {
		fmt.Println("我要上北大")

		if score > 690 {
			fmt.Println("我要学习厨师")
		} else if score > 685 {
			fmt.Println("我要学习盗墓")
		} else {
			fmt.Println("我要学习搬砖")
		}
	}
}
