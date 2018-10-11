package main

import "fmt"

type person struct {
	name string
	sex  string
	age  int
}

type student struct {
	//子结构体（类）继承父结构体（类）
	//通过匿名字段实现继承
	person

	score int
	id    int
	addr  string
}

type teacher struct {
	person//数据类型

	subject  string
	worktime int
}

func main() {

	var stu student = student{person{"菁菁", "女", 18}, 100, 102, "广东中山"}
	//var tea teacher=teacher{person{"兰骑士","女",24},"C语言",4}
	//stu.score = 100
	//stu.id = 101
	//stu.addr = "长达小区"
	//
	//stu.person.name = "静静"
	//stu.person.sex = "女"
	//stu.person.age = 18

	var tea teacher
	//子类可以直接使用父类的结构体成员信息
	tea.subject = "母猪的产后护理"
	tea.worktime = 4
	tea.name = "兰骑士"
	tea.age = 24
	tea.sex = "女"

	fmt.Println(stu)
	fmt.Println(tea)
}

/*
结构体嵌套结构体  继承关系实现
type 消费记录 struct{
	时间
	地点
	金额
	备注
	流水号
	消费类型
}
type 信用卡 struct{
	卡号
	额度
	密码
	类型
	手机

	消费记录	数组
}
 */
/*
type 技能 struct{
	名称
	范围
	伤害
	CD
	等级
	耗蓝
	buff
}
type 游戏角色 struct{
	职业
	性别
	等级
	经验
	血量
	蓝量
	物防
	魔防

	技能 数组

}
 */
