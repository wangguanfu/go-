package main

import "fmt"

// 使用append 完成字符串去空值。
func noEmpty(string []string) []string {
	// 创建一个空切片，用来存储 非空字符串
	out := string[:0]  // make([]string, 0)
	for _, word := range string {
		if word != "" {
			out = append(out, word)
		}
	}
	return out
}

// 不使用append 完成字符串去空值。
func noEmpty2(string []string) []string {
	i := 0
	for _, word := range string {
		if word != "" {
			string[i] = word
			i++
		}
	}
	return string[:i]
}

func main()  {
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	//afterData := noEmpty(data)
	afterData := noEmpty2(data)
	fmt.Println("data=", data)
	fmt.Println("afterData=", afterData)
}
