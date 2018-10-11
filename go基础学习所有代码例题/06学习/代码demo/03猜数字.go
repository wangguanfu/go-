package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	s := make([]int, 3)
	s[0] = rand.Intn(9) + 1 //1-9
	s[1] = rand.Intn(10)    //0-9
	s[2] = rand.Intn(10)    //0-9

	num := 0
	count := 0
	value := make([]int, 3)
	for {
		fmt.Println("请输入一个三位数")
		fmt.Scan(&num)
		if num <= 999 && num >= 100 {

			value[0] = num / 100     //百位
			value[1] = num / 10 % 10 //十位
			value[2] = num % 10      //个位

			for i := 0; i < 3; i++ {
				if s[i] == value[i] {
					count++
					fmt.Printf("第%d位相同\n", i+1)
				} else if s[i] > value[i] {
					fmt.Println("输入的数太小了")
				} else {
					fmt.Println("输入的数太大了")

				}
			}
			if count == 3 {
				fmt.Println("恭喜您，猜对了\n")
				break
			} else {
				count = 0
			}

		}
	}

}
