package main

import "fmt"

func test(i int) {
	var arr [10]int

	//优先使用错误拦截 在错误出现之前进行拦截 在错误出现后进行错误捕获
	//错误拦截必须配合defer使用  通过匿名函数使用
	defer func() {
		//恢复程序的控制权
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	arr[i] = 123 //err panic
	fmt.Println(arr)
}

func main() {

	i := 10

	test(i)

	fmt.Println("hello world")
}
