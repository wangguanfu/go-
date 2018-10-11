package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {

	f, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()

	n, err := f.WriteString("hello world\r\n")
	if err != nil {
		fmt.Println("WriteString err:", err)
		return
	}
	fmt.Println("n = ", n)

	//off, err := f.Seek(5, io.SeekStart)
	off, err := f.Seek(-5, io.SeekEnd)
	if err != nil {
		fmt.Println("Seek err:", err)
		return
	}
	fmt.Println("off:", off)

	n, err = f.WriteAt([]byte("1111"), off)
	if err != nil {
		fmt.Println("WriteAt err:", err)
		return
	}

	fmt.Println("write successful")
}
