package main

import "fmt"

func main0801() {

	m:=map[int]string{101:"刘备",103:"貂蝉",105:"李逵"}

	//对map中的值键判断 返回值为 value值 ok 键对应值是否存在
	value,ok:=m[101]

	if ok ==true{
		fmt.Println(value)
	}else{
		fmt.Println("未找到数据")
	}
}


func main0802(){
	m:=map[int]string{101:"刘备",103:"貂蝉",105:"李逵"}

	fmt.Println(m)
	//删除map元素 delete(map,键)
	delete(m,105)
	delete(m,101)
	fmt.Println(m)

}