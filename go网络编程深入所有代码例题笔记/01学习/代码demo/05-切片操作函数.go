package main

import "fmt"

/*func main()  {

	var s1 = []int {1, 2, 3}

	fmt.Println("s1=", s1, "cap=", cap(s1))

	s1 = append(s1, 10)

	fmt.Println("s1=", s1, "cap=", cap(s1))
	s1 = append(s1, 11)
	s1 = append(s1, 12)
	s1 = append(s1, 13)

	fmt.Println("s1=", s1, "cap=", cap(s1))

}*/
func main()  {
	var s1 = []int {1, 2, 3, 4, 5, 7}
	var s2 = []int {0, 0, 0, 0, 5, 7}

	copy(s2[1:4], s1[1:4])

	fmt.Println("s2=", s2)

}
