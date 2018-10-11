package main

import (
	"fmt"
)

func main()  {
	m1 := map[int]string{1: "Luffy", 2: "Sanji", 3:"Zoro", 5:"Nami"}

	m1[3] = "zhangsan"		//覆盖原有 map 中 key 值为 3 的 value
	for k, v := range m1 {
		fmt.Printf("[%d]%s\n", k, v)
	}

/*	for _, v := range m1 {		}  // 忽略 key 值，取 value
	for k，_ := range m1 {    }  // 忽略 value 值， 取 key 值*/

	// 默认 range 可以省略 value 值
	for k := range m1 {
		fmt.Printf("%v\n", k)
	}
	// 判断 map 中的某一个元素是否存在。
	if data, has := m1[5]; !has {	// has == false
		fmt.Printf("1, has=%v, data:%q\n", has, data)
	} else {
		fmt.Printf("2, has=%v, data:%q\n", has, data)
	}

}
