package main

import "fmt"

func main()  {

	// 使用数组方式创建切片
	var s1 = [] int {1, 2, 4, 6}
	fmt.Println("s1=", s1)

	fmt.Println("len=", len(s1), "cap=", cap(s1))

	// 使用make创建切片。 指定 len、指定 cap
	s2 := make([]string, 5, 10)		// 有空间，初值为 数据类型的默认值。

	fmt.Println("s2 = ", s2)
	fmt.Printf("s2[0] = %q\n", s2[0])
	fmt.Println("len=", len(s2), "cap=", cap(s2))

	// 使用make创建切片。 指定 len
	s3 := make([]int, 7)		// 没有指定容量，cap == len

	fmt.Println("s3 =", s3)
	fmt.Println("len=", len(s3), "cap=", cap(s3))

}
