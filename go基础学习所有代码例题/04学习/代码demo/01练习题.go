package main

import "fmt"

//     *		5	1	0*2+1
//    ***		4	3	1*2+1
//   *****		3	5	2*2+1
//  *******		2	7	3*2+1
// *********	1	9
//***********	0	11
func main() {
	//代码外层执行六次 外层控制行

	for i := 0; i < 6; i++ {
		//打印空格
		for j := 0; j < 6-i-1; j++ {
			fmt.Print(" ")
		}
		//打印星星
		for k := 0; k < i*2+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}
