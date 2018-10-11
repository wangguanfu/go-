package main

import "fmt"

/*func main()  {
	arr := [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("len(arr) = ", len(arr))
	fmt.Println("cap(arr) = ", cap(arr))

	s1 := arr[1:5:5]		// {2, 3, 4, 5}

	fmt.Println("s1 = ", s1)
	fmt.Println("len(s1) = ", len(s1))
	fmt.Println("cap(s1) = ", cap(s1))

	s2 := s1[2:3]			// {2, 3, 4, 5, 6, 7}  -- max : 4

	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2))
	fmt.Println("cap(s2) = ", cap(s2))
}*/

func main()  {
	arr := [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := arr[2:5]	// 3, 4, 5 --> 3 666 5
	s1[1] = 666
	fmt.Println("s1=", s1)

	s2 := s1[2:7]	// 5, 6, 7, 8, 9 --> 5, 6, 777, 8, 9
	s2[2] = 777
	fmt.Println("s2=", s2)
	fmt.Println("cap(s2) = ", cap(s2))

	fmt.Println("arr=", arr)
}
