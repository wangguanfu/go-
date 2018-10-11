package main

import "fmt"

//map作为函数参数
func test2(m map[int]string) {

	//修改map中的元素
	m[103] = "PDD"
	//添加数据
	m[104] = "段誉"

	delete(m, 102)
	fmt.Println(m)

}

func main0901() {

	//map中的键是单一数据类型 需要支持 == != 不支持复合类型 数组 切片
	m := map[int]string{101: "张飞", 102: "马菲菲", 103: "卢本伟"}
	//map作为函数参数  地址传递  形参可以改变实参的值
	test2(m)
	fmt.Println(m)

}

func main() {
	//map中的值为复合类型
	m := make(map[string][]int)

	m["毕向东"] = []int{60, 59, 61}
	m["尹志平"] = []int{99, 100, 101}
	m["黎活明"] = []int{10, 15, 20}

	for k, v := range m {
		//fmt.Println(k,v)
		sum := 0
		for _, data := range v {
			sum += data
		}
		fmt.Printf("%s 成绩为：%d  平均分：%d\n",k,sum,sum/len(v))

	}

}
