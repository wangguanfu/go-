package main

import "fmt"

func main() {
	var slice []interface{}

	slice = append(slice, 0, 3.14, [3]int{1, 1, 1}, "判断数据", []int{1, 2, 3}, 123, map[int]string{101: "蒙奇D路飞"}, 456)

	for _, v := range slice {
		//类型断言  判断数据的类型   只能判断接口数据类型
		//data,ok:=v.(map[int]string)
		//
		//if ok{
		//	fmt.Println(data)
		//}
		//断言格式  值,ok:=接口变量.(数据类型)
		if data, ok := v.(int); ok == true {
			fmt.Printf("整型数据：%d\n", data)
		} else if data, ok := v.(float64); ok {
			fmt.Printf("浮点型数据：%f\n", data)
		} else if data, ok := v.([]int); ok {
			fmt.Println("切片数据：", data)
		} else if data, ok := v.(map[int]string); ok {
			fmt.Println("map数据：", data)
		}

	}

}
