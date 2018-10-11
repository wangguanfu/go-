package main

import "fmt"

func main() {

	var str string
	fmt.Scan(&str)
	//字符切片
	ch := []byte(str)

	//统计字符个数
	var count [26]int

	for i := 0; i < len(ch); i++ {
		//计算字符和a偏移量（差值）
		offset := ch[i] - 'a'
		//count[offset]=count[offset]+1
		count[offset]++
	}
	//fmt.Println(count)

	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			fmt.Printf("字符：%c 出现%d次\n", 'a'+i, count[i])
		}
	}
}
