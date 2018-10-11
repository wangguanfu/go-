package main

import "fmt"

func main() {
	arr := [10]int{9, 1, 5, 6, 3, 7, 10, 8, 2, 4}

	//冒泡排序

	for i := 0; i < 10-1; i++ {
		for j := 0; j < 10-1-i; j++ {
			if arr[j] > arr[j+1] {

				//数据交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//temp := arr[j]
				//arr[j] = arr[j+1]
				//arr[j+1] = temp
			}
		}
		fmt.Println(arr)
	}

	//fmt.Println(arr)
}
