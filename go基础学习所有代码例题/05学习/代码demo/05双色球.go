package main

import (
	"math/rand"
	"time"
	"fmt"
)

//双色球 红球1-33 选择6个 不能重复  蓝球 1-16 选择1个
func main0501() {

	rand.Seed(time.Now().UnixNano())
	var red [6]int

	for i := 0; i < len(red); i++ {
		//red[i] = rand.Intn(33) + 1
		temp := rand.Intn(33) + 1
		//去重
		for j := 0; j < i; j++ {
			//重复
			if red[j] == temp {
				temp = rand.Intn(33) + 1
				j = -1
			}
		}
		red[i] = temp
	}

	fmt.Println(red, "+", rand.Intn(16)+1)

}
func main() {
	rand.Seed(time.Now().UnixNano())
	var red [6]int

	for i := 0; i < len(red); i++ {
		flag := false
		temp := rand.Intn(33) + 1

		for j := 0; j < i; j++ {
			//找到重复 跳出本层程序
			if red[j] == temp {
				flag = true
				i--
				break
			}
		}

		if flag == false {
			red[i] = temp
		}

	}
	fmt.Println(red, "+", rand.Intn(16)+1)
}
