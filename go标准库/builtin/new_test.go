package builtin_test

import (
	"fmt"
	"testing"
)


type man struct {
	Name string
	Age uint
}

// 创建基本类型
func TestBase(t *testing.T) {

}


// 创建对象
func TestObject(t *testing.T)  {

	// 分配内存 默认初始化 返回对象
	var m1 man
	// 分配内存 默认初始化 返回指针
	m2 := new(man)
	// 分配内存 显示初始化 返回对象
	m3 := man{}

	fmt.Printf("var m1 man:")
}

// 创建map
func TestArray(t *testing.T)  {

	var m1 map[string]string
	m2 := make(map[string]string,1)
	m3 := map[string]string{}

	m1["sb"] = "kang"
	m2["sb"] = "kang"
	m3["sb"] = "kang"
	fmt.Println(&m2)
	fmt.Println(&m3)


	var s1 []string
	var s4 [2]string
	s2 := make([]string,2)
	s3 := []string{}

	s1 =  append(s1, "sf")
	s4[0] = "sfsf"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}


// 创建slice
func TestSlice(t *testing.T)  {

	var m1 map[string]string
	m2 := make(map[string]string,1)
	m3 := map[string]string{}

	//m1["sb"] = "kang"
	m2["sb"] = "kang"
	m3["sb"] = "kang"
	fmt.Println(&m1)
	fmt.Println(&m2)
	fmt.Println(&m3)


	var s1 []string
	var s4 [2]string
	s2 := make([]string,2)
	s3 := []string{}

	s1 =  append(s1, "sf")
	s4[0] = "sfsf"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}


// 创建channel
func TestMap(t *testing.T)  {

	var m1 map[string]string
	m2 := make(map[string]string,1)
	m3 := map[string]string{}

	//m1["sb"] = "kang"
	m2["sb"] = "kang"
	m3["sb"] = "kang"
	fmt.Println(&m1)
	fmt.Println(&m2)
	fmt.Println(&m3)


	var s1 []string
	var s4 [2]string
	s2 := make([]string,2)
	s3 := []string{}

	s1 =  append(s1, "sf")
	s4[0] = "sfsf"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}




// 创建channel
func TestChannel(t *testing.T)  {

	var m1 map[string]string
	m2 := make(map[string]string,1)
	m3 := map[string]string{}

	//m1["sb"] = "kang"
	m2["sb"] = "kang"
	m3["sb"] = "kang"
	fmt.Println(&m1)
	fmt.Println(&m2)
	fmt.Println(&m3)


	var s1 []string
	var s4 [2]string
	s2 := make([]string,2)
	s3 := []string{}

	s1 =  append(s1, "sf")
	s4[0] = "sfsf"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

