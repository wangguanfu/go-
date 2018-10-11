package os__test

import (
	"fmt"
	"os"
	"testing"
)

// Hostname返回内核提供的主机名。
func Test_Hostame(t *testing.T) {
	fmt.Println(os.Hostname())
}

// Getpagesize返回底层的系统内存页的尺
func Test_GetPageSize(t *testing.T) {
	fmt.Println(os.Getpagesize())
}

// Environ返回表示环境变量的格式为"key=value"的字符串的切片拷贝。
func Test_Environ(t *testing.T) {
	fmt.Println(os.Environ())
}

// Getenv检索并返回名为key的环境变量的值。如果不存在该环境变量会返回空字符串。
func Test_Getenv(t *testing.T) {
	fmt.Println(os.Getenv("GOROOT"))
}

// Setenv设置名为key的环境变量。如果出错会返回该错误。
func Test_Setenv(t *testing.T) {
	fmt.Println(os.Setenv("my_name", "xmge"))
	fmt.Println(os.Getenv("my_name"))
}

// Clearenv删除所有环境变量。
func Test_Clarenv(t *testing.T) {
	os.Clearenv()
	fmt.Println(os.Environ())
}

// Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。
func Test_Exit(t *testing.T) {
	go func() {
		os.Exit(1)
	}()
	go func() {
		os.Exit(0)
	}()
}

// Expand用mapping 函数指定的规则替换字符串中的${var}或者$var（注：变量之前必须有$符号）。比如，os.ExpandEnv(s)等效于os.Expand(s, os.Getenv)。
func Test_Expand(t *testing.T) {
	mapping := func(key string) string {
		m := make(map[string]string)
		m = map[string]string{
			"world": "kitty",
			"hello": "hi",
		}
		if m[key] != "" {
			return m[key]
		}
		return key
	}
	s := "hello,world"            //  hello,world，由于hello world之前没有$符号，则无法利用map规则进行转换
	s1 := "$hello,$world $finish" //  hi,kitty finish，finish没有在map规则中，所以还是返回原来的值
	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))
}

// TODO 未测试成功
func Test_ExpandEnv(t *testing.T) {
	s := "hello $GOROOT"
	fmt.Println(os.ExpandEnv(s)) // hello /home/work/software/go，$GOROOT替换为环境变量的值，而h没有环境变量可以替换，返回空字符串

}

// Getuid返回调用者的用户ID。
func Test_Getuid(t *testing.T) {
	fmt.Println(os.Getuid())
}

//  Geteuid返回调用者的有效用户ID。
func Test_Geteuid(t *testing.T) {
	fmt.Println(os.Geteuid())
}

//  Getgroups返回调用者所属的所有用户组的组ID。
func Test_Getgid(t *testing.T) {
	fmt.Println(os.Getgid())
}

//  Getegid返回调用者的有效组ID。
func Test_Getgroups(t *testing.T) {
	fmt.Println(os.Getgroups())
}

// Getpid返回调用者所在进程的进程ID。
func Test_Getpid(t *testing.T) {
	fmt.Println(os.Getpid())
}

//  Getppid返回调用者所在进程的父进程的进程ID。
func Test_Getppid(t *testing.T) {
	fmt.Println(os.Getppid())
}
