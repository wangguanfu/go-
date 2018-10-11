package main

import (
	"os"
	"fmt"
)

/*func main()  {
	f, err := os.Create("C:/itcast/test.txt")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()

	fmt.Println("create successful")
}*/

/*func main()  {

	f, err := os.Open("C:/itcast/test.txt")
	if err != nil {
		fmt.Println("open err: ", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("hello world")
	if err != nil {
		fmt.Println("WriteString err: ", err)
		return
	}
	fmt.Println("open successful")
}*/

func main()  {
	f, err := os.OpenFile("C:/itcast/test.txt", os.O_RDWR | os.O_CREATE, 0600)  // 777--rwx rwx rwx
	if err != nil {
		fmt.Println("OpenFile err: ", err)
		return
	}
	defer f.Close()
	f.WriteString("hello world12345...")

	fmt.Println("open successful")
}
