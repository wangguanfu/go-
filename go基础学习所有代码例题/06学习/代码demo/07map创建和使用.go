package main

import "fmt"

func main0701() {

	//map[键数据类型]值数据类型

	//var m map[int]int=make(map[int]int,2)
	//
	//m[0]=123
	//m[100]=234
	//m[321]=10

	//键在map中是唯一的
	//链式存储  每次数据遍历是 不应定按照存储的顺序打印
	var m map[int]int = map[int]int{1: 123, 2: 321, 33: 12, 22: 12}

	//map数据格式会随着添加数据自动增长
	m[110] = 119
	//覆盖上一次的值
	m[110] = 120

	//fmt.Println(m)
	//打印map中的键和值
	//key  键  value  值
	//for k,v:=range m{
	//	fmt.Println(k,v)
	//}

	fmt.Println(len(m))
	fmt.Println(m)

}
func main0702() {
	m := make(map[string]int)

	m["hello"] = 1

	m["itcast"] = 2

	//	map中的键和值不能交换使用
	fmt.Println(m["hello"])

	for k, v := range m {
		fmt.Println(k, v)
	}

}

func main() {
	//统计字符出现次数
	var str string
	fmt.Scan(&str)

	//将字符串转成字符切片
	ch := []byte(str)

	m:=make(map[byte]int)
	for i := 0; i < len(ch); i++ {
		//fmt.Println(ch[i])
		m[ch[i]]++
	}

	for k,v:=range m {
		fmt.Printf("%c   %d\n",k,v)
	}

}
