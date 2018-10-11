package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main0401() {

	//1、导入头文件 math/rand  time
	//2、添加随机数种子  rand.Seed(time.Now().UnixNano())
	//3、使用随机数 rand,Intn(10)
	//当前系统时间
	//fmt.Println(time.Now())
	//用当前时间做随机数混淆
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		//默认随机数使用的时间为1970.1.1.0.0.0
		fmt.Println(rand.Intn(100)) //随机0-99之间的数   % 100

	}
}

func main0402() {

	rand.Seed(time.Now().UnixNano())
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	for i := 0; i < 10-1; i++ {
		for j := 0; j < 10-1-i; j++ {
			if arr[j] > arr[j+1] {

				//数据交换
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}

	fmt.Println(arr)

}

//双色球 红球 1-33 个数 不能重复 选择6个  篮球 1-16 选择1个

func main() {

	rand.Seed(time.Now().UnixNano())
	var redball [6]int

	for i := 0; i < len(redball); i++ {
		//redball[i]=rand.Intn(33)+1//1-33
		temp := rand.Intn(33) + 1

		//去重
		for j := 0; j < i; j++ {
			if redball[j] == temp {
				j = -1
				temp = rand.Intn(33) + 1
			}
		}
		redball[i] = temp
	}

	fmt.Println(redball, "+", rand.Intn(16)+1)

}
