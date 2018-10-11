package main

import "fmt"

func main0801() {
	var scores [10]int

	//通过键盘输入学生成绩 计算总分和平均分
	//fmt.Scan(&scores)//err
	for i := 0; i < len(scores); i++ {
		fmt.Scan(&scores[i])
	}
	sum := 0
	for _, v := range scores {
		sum += v
	}

	fmt.Printf("总成绩为：%d,平均成绩为：%d\n", sum, sum/len(scores))
	fmt.Println(scores)

}
func main() {
	//十只小猪称体重

	var weight [10]int

	for i, _ := range weight {
		fmt.Scan(&weight[i])
	}

	max := weight[0]
	for i := 1; i < len(weight); i++ {
		if weight[i] > max {
			max = weight[i]
		}
	}

	fmt.Println(max)

}
