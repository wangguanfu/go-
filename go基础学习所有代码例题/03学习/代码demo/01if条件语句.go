package main

import "fmt"

func main0101() {
	var score int

	fmt.Scan(&score)
	//if 表达式 {代码}  如果表达式为真 执行{}内的代码
	if score > 700 {
		fmt.Println("我要上清华")
	}else{
		fmt.Println("我要上浙大")
	}
}
func main0102() {
	var score int

	fmt.Scan(&score)
	//if 表达式1 {代码1}  如果表达式为真 执行{}内的代码
	//else if 表达式2 {代码2}
	//else {代码3}
	if score > 700 {
		fmt.Println("我要上清华")
	}else if score >680  {
		fmt.Println("我要上北大")
	}else{
		fmt.Println("我要上浙大")
	}
}
