package main

import (
	"strings"
	"fmt"
)

func wordCountFunc2(str string) map[string]int {
	// 按空格拆分字符串
	words := strings.Fields(str)
	m := make(map[string]int)		//创建初始容量为 0 的map

	for _, data := range words {
		m[data]++
	}
	return m
}

func wordCountFunc(str string) map[string]int {
	// 按空格拆分字符串
	words := strings.Fields(str)
	m := make(map[string]int)		//创建初始容量为 0 的map

	for _, data := range words {
		if _, ok := m[data]; !ok {	//  当 ok == false ，当前if 满足
			m[data] = 1
		} else {					// 当 ok == true , 满足
			m[data]++
		}
	}
	return m
}

func main()  {
	string := "I love my work and I love my family family family family too love love love"

	//retMap := wordCountFunc(string)
	retMap := wordCountFunc2(string)

	for k, v := range retMap {
		fmt.Printf("%s:%d\n", k, v)
	}

}
