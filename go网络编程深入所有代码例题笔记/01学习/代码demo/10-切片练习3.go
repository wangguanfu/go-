package main

import "fmt"

func remove(data []int, idx int) []int {

	copy(data[idx:], data[idx+1:])

	return data[:len(data)-1]
}

func main()  {

	data := []int {5, 6, 7, 8, 9}  // {5, 6, 7, 8, 9}  --> {5, 6, 8, 9, 9}

	afterData := remove(data, 3)

	fmt.Println("afterData:", afterData)
}
