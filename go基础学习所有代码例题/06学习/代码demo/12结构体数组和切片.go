package main

import "fmt"

type stu struct {
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main1201() {

	//var 结构体变量名 结构体数据类
	//var 结构体数组名 [元素个数]结构体数据类型

	var arr [3] stu = [3]stu{
		stu{101, "边振龙", 10, "男", 3, "河南道口"},
		stu{102, "段永超", 9, "男", 2, "北京五道口"},
		stu{102, "郑强", 9, "男", 2, "河南驻马店"}}
	//数组名[下标].成员
	//arr[0].id = 101
	//arr[0].name = "边振龙"
	//arr[0].sex = "男"
	//arr[0].score = 3
	//arr[0].addr = "河南道口"
	//arr[0].age = 10
	//
	//arr[1].id = 102
	//arr[1].name = "段永超"
	//arr[1].sex = "男"
	//arr[1].score = 2
	//arr[1].addr = "北京五道口"
	//arr[1].age = 9

	//fmt.Println(arr)
	//fmt.Println(len(arr))

	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main1202() {

	//结构体切片  内存地址为0
	var slice []stu = make([]stu, 1)

	slice[0].id = 101
	slice[0].name = "边振龙"
	slice[0].sex = "男"
	slice[0].score = 3
	slice[0].addr = "河南道口"
	slice[0].age = 10

	//slice[1].id = 101
	//slice[1].name = "边振龙"
	//slice[1].sex = "男"
	//slice[1].score = 3
	//slice[1].addr = "河南道口"
	//slice[1].age = 10

	slice = append(slice, stu{102, "段永超", 9, "男", 2, "北京五道口"},
		stu{102, "郑强", 9, "男", 2, "河南驻马店"})

	fmt.Println(slice)
}

func main() {
	//结构体切片
	s := []stu{stu{102, "段永超", 9, "男", 2, "北京五道口"},
		stu{101, "郑强", 9, "男", 100, "河南驻马店"},
		stu{101, "边振龙", 10, "男", 3, "河南道口"}}

	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			//根据成绩进行排序
			if s[j].score > s[j+1].score {
				//交换信息   交换结构体切片元素
				s[j], s[j+1] = s[j+1], s[j]
				//结构体切片里面的成员允许交换
				//s[j].id,s[j+1].id=s[j+1].id,s[j].id

			}
		}
	}

	for i:=0;i<len(s);i++{
		fmt.Println(s[i])
	}
}
