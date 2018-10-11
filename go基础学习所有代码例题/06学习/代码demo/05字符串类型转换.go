package main

func main0501() {

	//将bool转成字符串
	//str:=strconv.FormatBool(false)
	//fmt.Println(str)

	//FormatInt(数据，进制)
	//str:=strconv.FormatInt(123,10)//2-36  能够使用的是2-36 （0-9 a-z）
	//fmt.Println(str)
	//将十进制整型转成字符串
	//str:=strconv.Itoa(666)
	//fmt.Println(str)

	//将浮点型转成字符串  保留具体位数的小数 会四舍五入
	//str:=strconv.FormatFloat(3.1415926,'f',4,64)
	//fmt.Println(str)
}


func main(){

	////将字符串转成bool
	//b,err:=strconv.ParseBool("true")
	//if err!=nil{
	//	fmt.Println("出错")
	//}else{
	//
	//	fmt.Println(b)
	//}
	//将字符串转成整型
	//i,err:=strconv.ParseInt("1111011",2,64)
	//if err!=nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Println(i)
	//
	//}
	//将字符串转成十进制整型
	//i,_:=strconv.Atoi("123abc")
	//fmt.Println(i)
	//fmt.Printf("%T\n",i)

	//将字符串转成float类型
	//f,_:=strconv.ParseFloat("3.14159",64)
	//fmt.Println(f)
}