package builtin_test

import (
	"fmt"
	"testing"
	"time"
)

// 获取复数的实数部分
func TestReal(t *testing.T) {
	var f = complex(1, 2)
	fmt.Println(real(f))
}

// 获取复数的虚数部分
// 翻译： 虚数[imaginary number][iˈmædʒinəri ˈnʌmbə]
func TestImag(t *testing.T) {
	var f = complex(1, 2)
	fmt.Println(imag(f))
}

// 生成一个复数
func TestComplex(t *testing.T) {
	var f complex64 = complex(1, 2)
	fmt.Print(f)
}

// 创建结构体对象 初始化并返回对象指针
func TestNew(t *testing.T) {
	var f *complex64 = new(complex64) // 这种定义方式可以提高代码可读性
	fmt.Print(f)
}

// 定义 slice、map、channel
func TestMake(t *testing.T) {
	var s = make([]string, 10, 10)
	s = append(s, "1")
	var m = make(map[string]interface{}, 1)
	var c = make(chan int)
	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(c)
}

// The cap built-in function returns the capacity of v, according to its type:
//	Array: the number of elements in v (same as len(v)).
//	Pointer to array: the number of elements in *v (same as len(v)).
//	Slice: the maximum length the slice can reach when resliced;
//	if v is nil, cap(v) is zero.
//	Channel: the channel buffer capacity, in units of elements;
//	if v is nil, cap(v) is zero.
func TestCap(t *testing.T) {
	var s = make([]string, 1, 10)
	fmt.Printf("cap:%d\n", cap(s))
}

// The len built-in function returns the length of v, according to its type:
//	Array: the number of elements in v.
//	Pointer to array: the number of elements in *v (even if v is nil).
//	Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
//	String: the number of bytes in v.
//	Channel: the number of elements queued (unread) in the channel buffer;
//	if v is nil, len(v) is zero.
func TestLen(t *testing.T) {
	var s = make([]string, 1, 10)
	fmt.Printf("len:%d\n", len(s))
}

// The append built-in function appends elements to the end of a slice. If
// it has sufficient capacity, the destination is resliced to accommodate the
// new elements. If it does not, a new underlying array will be allocated.
// Append returns the updated slice. It is therefore necessary to store the
// result of append, often in the variable holding the slice itself:
//	slice = append(slice, elem1, elem2)
//	slice = append(slice, anotherSlice...)
// As a special case, it is legal to append a string to a byte slice, like this:
//	slice = append([]byte("hello "), "world"...)
func TestAppend(t *testing.T) {
	var s = make([]int, 4, 10)
	s = append(s, 1)
	fmt.Println(s)
}

// The copy built-in function copies elements from a source slice into a
// destination slice. (As a special case, it also will copy bytes from a
// string to a slice of bytes.) The source and destination may overlap. Copy
// returns the number of elements copied, which will be the minimum of
// len(src) and len(dst).
func TestCopy(t *testing.T) {
	s1 := []string{}
	s2 := []string{"x", "m", "g", "e"}
	copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)
}

// The delete built-in function deletes the element with the specified key
// (m[key]) from the map. If m is nil or there is no such element, delete
// is a no-op.
func TestDelete(t *testing.T) {
	m := make(map[string]interface{}, 1)
	delete(m, "1")
}


// 内建函数close关闭信道，该通道必须为双向的或只发送的。它应当只由发送者执行，
// 而不应由接收者执行，其效果是在最后发送的值被接收后停止该通道。在最后的值从已关闭的信道中被接收后，
// 任何对其的接收操作都会无阻塞的成功。对于已关闭的信道，语句：
// x, ok := <-c
// 还会将ok置为false。
func TestClose(t *testing.T) {
	var c chan int     // 不会分配空间
	c = make(chan int) // 分配空间
	close(c)           // 关闭通道
}


// The panic built-in function stops normal execution of the current
// goroutine. When a function F calls panic, normal execution of F stops
// immediately. Any functions whose execution was deferred by F are run in
// the usual way, and then F returns to its caller. To the caller G, the
// invocation of F then behaves like a call to panic, terminating G's
// execution and running any deferred functions. This continues until all
// functions in the executing goroutine have stopped, in reverse order. At
// that point, the program is terminated and the error condition is reported,
// including the value of the argument to panic. This termination sequence
// is called panicking and can be controlled by the built-in function
// recover.
func TestPanic(t *testing.T) {
	wait := make(chan bool)
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("测试panic:[%d]\n", i)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		time.Sleep(5 * time.Second)
		panic("5s后panic")
	}()
	<-wait
}


// The recover built-in function allows a program to manage behavior of a
// panicking goroutine. Executing a call to recover inside a deferred
// function (but not any function called by it) stops the panicking sequence
// by restoring normal execution and retrieves the error value passed to the
// call of panic. If recover is called outside the deferred function it will
// not stop a panicking sequence. In this case, or when the goroutine is not
// panicking, or if the argument supplied to panic was nil, recover returns
// nil. Thus the return value from recover reports whether the goroutine is
// panicking.
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("遇到重大故障")
}


// The println built-in function formats its arguments in an
// implementation-specific way and writes the result to standard error.
// Spaces are always added between arguments and a newline is appended.
// Println is useful for bootstrapping and debugging; it is not guaranteed
// to stay in the language.
func TestPrint(t *testing.T) {
	print("hello error")
}


// The println built-in function formats its arguments in an
// implementation-specific way and writes the result to standard error.
// Spaces are always added between arguments and a newline is appended.
// Println is useful for bootstrapping and debugging; it is not guaranteed
// to stay in the language.
func TestPrintln(t *testing.T)  {
	println("hello error")
}


