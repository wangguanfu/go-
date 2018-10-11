package main

import (
	"fmt"
	"unsafe"
)

type Person2 struct {
	name string
	sex byte
	age int
}

func main()  {
/*	// 1 顺序初始化
	var p1 *Person2 = &Person2{"n1", 'f', 11}
	fmt.Println("p=", p1)

	// 2 指定成员初始化
	p2 := &Person2{name:"n1", sex:'f'}
	fmt.Println("p2=", p2)

	// 3. 使用另一个结构体地址初始
	var p3 *Person2
	p3 = p2
	fmt.Println("p3=", p3)

	// 4. new, 指定成员赋值
	p4 := new(Person2)
	fmt.Println("p4=", p4)
	p4.name = "n4"
	p4.age = 119
	p4.sex = 'm'
	fmt.Println("p4=", p4)

	// 指针变量的值(地址) == 结构体首元素的地址
	fmt.Printf("p4 = %p\n", p4)
	fmt.Printf("p4.name = %p\n", &p4.name)

	fmt.Println("main 1: p4=", p4)
	test(p4)
	fmt.Println("main 2: p4=", p4)*/

	var temp Person2
	fmt.Println("main 1: temp=", temp, "temp size:", unsafe.Sizeof(temp))

	test(&temp)

	fmt.Println("main 2: temp=", temp)
}

func test(p *Person2) {
	p.name = "111111"
	p.sex = 'N'
	p.age = 111

	fmt.Println("test: p=", p,  "temp size:", unsafe.Sizeof(p))
}