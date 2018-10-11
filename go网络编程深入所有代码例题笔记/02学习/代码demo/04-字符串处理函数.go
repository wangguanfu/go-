package main

import (
	"strings"
	"fmt"
)

/*func main()  {
	flg := strings.Contains("hello", "lloi")
	fmt.Println("flg = ", flg)

	str := []string {"hello", "go", "haha", "xixi", "hoho"}
	retStr := strings.Join(str, "%")
	fmt.Println("retStr:", retStr)

	res := strings.Trim("hello", "o")
	fmt.Printf("res:%q\n", res)

	data := "a man a plan a canal panama a"
	resStr := strings.Replace(data, "a", "abc", 2)
	fmt.Printf("res:%q\n", resStr)
}*/

func main()  {

	str := "I love my work and I love my family too"
	/*	ret := strings.Split(str, "I ")
		for _, word := range ret {
			fmt.Println(word)
		}*/
	retStr := strings.Fields(str)
	for _, word := range retStr {
		fmt.Println(word)
	}

	flg := strings.HasSuffix("word cup.png", ".jpg")
	fmt.Println("flg = ", flg)

	flg =strings.HasPrefix("word cup.png", "worl")
	fmt.Println("flg = ", flg)

}
