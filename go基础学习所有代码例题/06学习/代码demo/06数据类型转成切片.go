package main

import (
	"strconv"
	"fmt"
)

func main() {
	a:=10
	b:=true
	c:=3.14159
	d:="hello"

	slice:=make([]byte,0,1024)

	slice=strconv.AppendBool(slice,b)
	slice=strconv.AppendInt(slice,int64(a),10)
	slice=strconv.AppendFloat(slice,c,'f',3,64)
	slice=strconv.AppendQuote(slice,d)

	fmt.Printf("%s\n",slice)

}
