package main

import "fmt"

func noSame(str []string) []string {
	out := str[:1]		// "red" --》 out

	for _, word := range str {
		i := 0
		for ; i < len(out); i++ {
			if word == out[i] {
				break
			}
		}
		if i == len(out) {
			out = append(out, word)
		}
	}
	return out
}

func main()  {
	color := []string{"red", "black","Green", "red", "Green","Green","Green","Green", "pink", "blue", "pink", "blue"}
	data := noSame(color)
	fmt.Println("data:", data)
}
