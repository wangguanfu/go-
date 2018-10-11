package main

import "fmt"

func main() {

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//定义数组下标
	i := 0
	j := len(arr) - 1

	for i <= j {
		//if i>j {
		//	break
		//}
		//数据交换
		arr[i], arr[j] = arr[j], arr[i]

		i++
		j--

	}
	fmt.Println(arr)
}
