package main

import "fmt"

//记者：我是记者  我的爱好是偷拍 我的年龄是34 我是一个男狗仔
//程序员：我叫孙全 我的年龄是23 我是男生 我的工作年限是 3年

type person7 struct {
	name string
	age  int
	sex  string
}

//记者子类
type rep struct {
	person7
	hooby string
}

//程序员子类
type dep struct {
	person7
	worktime int
}

func (p *person7) SayHi() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我是%s生，", p.name, p.age, p.sex)
}

func (r *rep) SayHello() {
	r.SayHi()
	fmt.Printf("我的爱好是%s\n", r.hooby)
}

func (d *dep) SayHello() {
	d.SayHi()
	fmt.Printf("我的工作时间是%d年\n", d.worktime)
}

func main() {

	var r rep = rep{person7{"卓伟", 40, "男"}, "偷拍"}

	var d dep = dep{person7{"图灵", 28, "男"}, 4}

	//根据数据类型调用不同的方法
	r.SayHello()
	d.SayHello()

}
