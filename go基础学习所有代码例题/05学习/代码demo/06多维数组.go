package main

import "fmt"

func main0601() {
	//一维数组
	//var 数组名 [元素个数]数据类型

	//二维数组
	//var 数组名 [行][列]数据类型
	//数组元素个数=行*列

	//var arr [3][4]int = [3][4]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}}
	//数组如果初始化部分值  未初始化值为0
	//var arr [3][4]int = [3][4]int{{1, 2, 3, 4}, {3}}
	//数组可以找到具体下标可以赋值
	//var arr[3][4]int=[3][4]int{1:{1,2,3,4},2:{3,4,5,6},0:{'a','b','c','d'}}

	arr := [3][4]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}}

	//arr[1][1]=10

	//fmt.Println(arr)

	//fmt.Println(len(arr)) //len(二维数组)二维数组行个数
	//arr[0] 表示整个一维数组
	//fmt.Println(len(arr[0])) //len(位数数组[下标])二维数组列个数
	//fmt.Println(arr[0])

	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr[0]); j++ {
	//		fmt.Print(arr[i][j]," ")
	//	}
	//	fmt.Println()
	//}
	//c 表示行的下标 v 二维数组中一行的值
	for _, v := range arr {
		//fmt.Println(c,v)
		for _, data := range v {
			//fmt.Printf("行：%d  列：%d  值：%d\t",c,l,data)
			fmt.Print(data)
		}
		fmt.Println()
	}
}
func main() {
	//var 数组名 [元素个数]数据类型
	//var 数组名 [行][列]数据类型
	//三维数组 var 数组名 [层][行][列] 数据类型
	var arr [2][3][4]int = [2][3][4]int{
		{
			{1, 2, 3, 4},
			{2, 3, 4, 5},
			{3, 4, 5, 6}},
		{
			{4, 5, 6, 7},
			{5, 6, 7, 8},
			{6, 7, 8, 9}}}

	fmt.Println(arr)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				fmt.Print(arr[i][j][k])
			}
			fmt.Println()
		}
	}
	//四维数组
	//var a [2][3][4][5]int
	//五维数组
	//var b [2][3][4][5][6]int
}
