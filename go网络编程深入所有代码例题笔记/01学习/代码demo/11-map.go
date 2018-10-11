package main

import "fmt"

func main()  {
	// 1.
	var m1 = map[int]string{1:"red", 5021:"black", 502:"yellow"}
	fmt.Println("m1=", m1)

	// 2.
	var m2 = map[int]string{}
	m2[4] = "abc2"
	m2[4] = "abc8910"		// key 与已有key 重复，会直接覆盖原有key -- value
	fmt.Println("m2=", m2)

	// 3.
	m3 := make(map[int]string, 13)		// 0： 表示初始容量 。 ——map 没有 cap
	m3[45] = "abc"
	m3[46] = "abc"
	m3[47] = "abc"
	m3[452] = "abc"

	fmt.Println("m3=", m3, "len=", len(m3)/*, "cap=", cap(m3)*/)

	m4 := make(map[string]string)			// 【常用】  map 可以动态扩容。

	m4["a"] = "aaaa"
	m4["b"] = "bbbb"
	m4["c"] = "cccc"
	m4["d"] = "dddd"

	fmt.Println("m4=", m4)

}