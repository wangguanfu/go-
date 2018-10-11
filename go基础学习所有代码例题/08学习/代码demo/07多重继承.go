package main

import "fmt"

type human1 struct {
	name string
	sex  string
}
type person5 struct {
	id  int
	age int
}

type student5 struct {
	human1
	person5

	score int
	addr  string
}

func main() {


	stu := student5{human1{"秦海东", "男"}, person5{101, 20}, 101, "江苏启东"}
	fmt.Println(stu)

}
