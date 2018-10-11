package main

import (
	"strings"
	"fmt"
)

func main0401() {

	//str := "hello world"
	//
	//s := "hehe"

	str := "传智播客"
	s := "智播"
	//查找一个字符串是否出现在另外一个字符串中 返回值是bool类型
	//strings.Contains 可以用于模糊查找
	b := strings.Contains(str, s)
	if b == true {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}

func main0402() {
	//str:=[]string{"1234","2234","3324","1238"}
	str := []string{"传", "智", "播", "客"}
	//字符串拼接  可以将一个字符串切片组合成一个字符串，可以使用连接符
	s := strings.Join(str, "-")

	fmt.Println(s)
}

func main0403() {
	str := "hello worwoldwo"
	s := "wo"

	sum := 0 //计算下标整体偏移值
	//查找一个字符串在另一个字符串中第一次出现的位置 返回值为整型 数据下标 如果没找到返回值为-1
	i := strings.Index(str, s)
	fmt.Println(i)

	//将字符串部分转成字符切片
	slice := []byte(str[i+len(s):])
	//fmt.Printf("%s\n",slice)
	sum += i + len(s)

	temp := string(slice)

	i = strings.Index(temp, s)
	fmt.Println(sum + i)


}

func main0404(){
	str:="hello"
	//将一个字符串重复多次
	s:=strings.Repeat(str,2)
	fmt.Println(s)
}


func main0405(){
	//str:="hello world"
	//
	////将一个字符串中的字符串用另外一个字符串替换  替换n次  -1表示全部替换
	//s:=strings.Replace(str,"l","m",-1)
	//
	//fmt.Println(s)

	str:="我x你大爷，你大爷是个傻x"

	s:=strings.Replace(str,"大爷","**",-1)
	fmt.Println(s)

}

func main0406(){

	//str:="1234-2234-3234-4562"
	//str:="www.itcast.cn"
	str:="123456@qq.com"
	//s:=strings.Split(str,"-")
	//s:=strings.Split(str,".")
	s:=strings.Split(str,"@")

	fmt.Println(s)
}

func main0407(){
	//str:="   a u ok?    "
	str:="===a=u=ok?====="

	//去掉头尾指定字符串
	//s:=strings.Trim(str," ")
	s:=strings.Trim(str,"=")

	fmt.Printf("%s",s)

}

func main0408(){

	str:="   a u ok?    "
	//去掉字符串中的空格 并转成切片类型
	s:=strings.Fields(str)

	fmt.Println(s)
	fmt.Println(len(s))
}