package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mapDelete(m map[int]string, key int)  {

	for i:=10; i<=50; i++ {
		m[i] = "aaa" + strconv.Itoa(i)
	}
	fmt.Println("mapDelete: m = ", m, "len = ", len(m))

	delete(m, key)

	strings.Fields()
}

func main()  {
	m1 := map[int]string{1: "Luffy", 2: "Sanji", 3:"Zoro", 5:"Nami"}

	mapDelete(m1, 3)

	fmt.Println("main : m1=", m1, "len = ", len(m1))
}
