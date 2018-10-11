package main

import (
	"fmt"
)

func swap(a, b int)  {
	a, b = b, a
	fmt.Printf("swap: a=%d, b=%d\n", a, b)
}

func swap2(a, b *int)  {
	*a, *b = *b, *a
	fmt.Printf("swap2: *a=%d, *b=%d\n", *a, *b)
}

func main()  {
	var a int = 10
	var b int = 300

	fmt.Printf("main 1: a=%d, b=%d\n", a, b)
	//swap(a, b)
	swap2(&a, &b)
	fmt.Printf("main 2: a=%d, b=%d\n", a, b)

}