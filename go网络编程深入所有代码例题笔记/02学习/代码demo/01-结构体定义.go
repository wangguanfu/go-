package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	sex byte
	age int
}

type student struct {
	p Person			// 结构体嵌套
	score [1000]int
}

func initFunc(p Person) Person {
	p.name = "name1"
	p.age = 19
	p.sex = 'f'

	fmt.Println("initFunc p:", p)
	fmt.Println("tmp size:", unsafe.Sizeof(p))

	return p
}


func main()  {
	// 1. 顺序初始化, 必须初始化所有的成员变量
	var p1 Person = Person{"Luffy", 'm', 99}
	fmt.Println("p1=", p1)

	// 2. 指定成员初始化, 未初始化成员取默认值
	p2 := Person{name:"Zoro", age:76}
	fmt.Println("p2=",p2)

	// 取成员变量， 使用“.”
	p2.sex = 'f'
	p2.age = 111
	fmt.Println("p2=",p2)

	// 结构体地址 == 结构体首元素地址。
	fmt.Printf("&p2 = %p\n", &p2)
	fmt.Printf("&p2.name = %p\n", &p2.name)
	fmt.Printf("&p2.name = %p\n", &p2.sex)

	p3 := Person{"Luffy", 'm', 99}
	p4 := Person{"Luffy", 'm', 99}
	p5 := Person{"Luffy", 'm', 912}

	fmt.Println("p3 == p4 ?", p3 == p4)
	fmt.Println("p3 == p5 ?", p3 != p5)

	var tmp Person
	fmt.Println("tmp=", tmp)

	fmt.Println("tmp size:", unsafe.Sizeof(tmp))
	tmp = initFunc(tmp)

	fmt.Println("main tmp:", tmp)

}
