package main

import "fmt"

func main0401() {

	var m int

	fmt.Scan(&m)

	//switch 变量  case 值1:代码1 case 值2: 代码2 default:代码3
	switch m {
	case 1:
		fmt.Println("一月份")
	case 2:
		fmt.Println("二月份")
	default:
		fmt.Println("未知月份")
	}
}
func main0402() {
	var score int
	fmt.Scan(&score)

	switch score / 10 {
	case 10:
		//fmt.Println("A")
		//让当前case向下执行
		fallthrough
	case 9:
		fmt.Println("A")
		fallthrough
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
		//fallthrough //err 不能在最后一个case中写fallthrough
	}
}
func main0403() {

	a := 10
	b := 20

	//if 优点可以判断区间 可以进行嵌套使用  缺点 执行效率比较低
	//if a > b {
	//	fmt.Println("a")
	//} else {
	//	fmt.Println("b")
	//}

	//switch 语句 执行效率高  缺点 不能判断复杂区间 不能嵌套使用
	switch a > b {

	case true:
		fmt.Println(a)
	case false:
		fmt.Println(b)

	}
}

func main() {
	//fmt.Println("???")

	var value float64=3.14
	//fmt.Printf("%.20f",value)

	//fmt.Scan(&value)
	//可以作为switch 的参数使用  不建议使用
	switch value {

	case 3.14000000000000012434:
		fmt.Println(3.14)
	//case 3.14:
		//fmt.Println(1.2)

	}
}
func main0405() {
	var score int
	fmt.Scan(&score)

	//switch 中如果有多个选项的值重复执行相同的代码  可以放在一起 中间用逗号分隔
	switch score / 10 {
	case 9, 10, 'A':
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
		//fallthrough //err 不能在最后一个case中写fallthrough
	}
}
func main0406() {
	var score int
	fmt.Scan(&score)
	//switch 可以在不写 参数的时候可以判断多个区间 执行效率和if相同
	switch {
	case score > 90:
		fmt.Println("A")
	case score > 80:
		fmt.Println("B")
	case score > 70:
		fmt.Println("C")
	case score > 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
		//fallthrough //err 不能在最后一个case中写fallthrough
	}
}
