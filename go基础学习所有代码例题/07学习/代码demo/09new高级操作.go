package main

import "fmt"

type Stu struct {
	id   int
	name string
}

func main0901() {

	//数组指针
	var p *[10]int

	p = new([10]int)

	//当new([10]int)时会在内存中创建一个连续的整型空间10个 默认值为0  数组指针变量p指向该内存
	for i := 0; i < len(p); i++ {
		p[i] = i
	}
	//fmt.Println(*p)

	for i, v := range p {
		fmt.Println(i, v)
	}

}

func main0902() {
	var p *[]int
	//通过new创建一个切片
	p = new([]int)

	fmt.Printf("%p\n", p)
	*p = append(*p, 1, 2, 3, 4, 5)
	*p = append(*p, 1, 2, 3)
	(*p)[1] = 222

	fmt.Printf("%p\n", p)
	fmt.Println(*p)

	//可以打印切片的长度和容量
	fmt.Println(len(*p))
	fmt.Println(cap(*p))

}

func main() {

	var p *Stu
	//创建一个结构体变量大小的空间
	p = new(Stu)

	fmt.Printf("%p\n", p)
	//p.id=101
	//p.name="郑强"
	*p = Stu{101, "郑强"}
	fmt.Printf("%p\n", p)

	fmt.Println(*p)
}
