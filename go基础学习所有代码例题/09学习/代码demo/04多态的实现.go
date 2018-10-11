package main

import "fmt"

//接口定义
type CSer interface {
	Socket()
	Jiami()
}

type framework struct {
	csckimp string //厂商
	data    string //数据
}

type ck1 struct {
	framework
}

type ck2 struct {
	framework
}

func (c1 *ck1) Socket() {
	fmt.Printf("%s正在socket连接中，数据传递%s\n", c1.csckimp, c1.data)
}

func (c1 *ck1) Jiami() {
	fmt.Println("数据加密\n")
}

func (c2 *ck2) Socket() {
	fmt.Printf("%s正在socket连接中，数据传递%s\n", c2.csckimp, c2.data)
}

func (c2 *ck2) Jiami() {
	fmt.Println("数据加密\n")
}


//定义一个普通函数  用接口作为函数参数
//多态
func Connect (i CSer){
	i.Socket()
	i.Jiami()
}
func main() {

	var c1 ck1 = ck1{framework{"东软国际","联网数据"}}

	//var i CSer
	//i=&c1

	Connect(&c1)
	//i.Socket()
	//i.Jiami()



	var c2 ck2=ck2{framework{"中软国际","访问数据库数据"}}
	//i=&c2

	Connect(&c2)
	//i.Socket()
	//i.Jiami()



}
