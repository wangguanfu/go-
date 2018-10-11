package main

import (
	"fmt"
	"errors"
)

func calc(a int, b int) (v int, err error) {

	//捕获错误信息
	if b == 0 {
		//如果代码中出现错误 可以使用errors.New()创建错误信息
		err = errors.New("除数不能为0")
		return
	}
	v = a / b
	return

}

func main() {

	a := 10
	b := 0
	v, err := calc(a, b)
	//根据错误信息进行处理
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	//fmt.Println(v)
}
