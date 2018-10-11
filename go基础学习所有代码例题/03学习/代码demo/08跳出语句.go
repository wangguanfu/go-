package main

import "fmt"

func main0801() {

	//死循环
	for{
		fmt.Println("跳出本层循环")

		//在循环中 跳出本层循环
		break
	}

}
func main0802(){
	//var name string
	//var passwd string
	//
	////为死循环添加break出口
	//for{
	//	fmt.Println("请输入用户名和密码")
	//	fmt.Scan(&name,&passwd)
	//
	//	if name=="admin" && passwd=="123456"{
	//		fmt.Println("成功")
	//		//跳出循环
	//		break
	//	}else{
	//		fmt.Println("输入错误")
	//	}
	//}
	sum:=0
	i:=0
	for  {
		sum+=i
		if i==100{
			break
		}
		i++
	}

	fmt.Println(sum)
}

func main0803(){

	sum:=0
	//1-100偶数的和
	for i:=1;i<=100;i++{
		if i%2==1{
			//结束本次循环 继续下次循环
			continue
		}
		sum+=i
	}

	fmt.Println(sum)
}

func main0804(){

	fmt.Println("hello1")
	fmt.Println("hello2")
	//goto 从一个标志位跳到另外一个标志位
	//goto用法很灵活可以在一个循环 函数 跳到另外一个循环 函数
	//goto FLAG;
	fmt.Println("hello3")
	fmt.Println("hello4")
	fmt.Println("hello5")
	//FLAG:
	fmt.Println("hello6")
	fmt.Println("hello7")




}

func main(){

	FLAG:
	fmt.Println("hello")

	goto FLAG;
}