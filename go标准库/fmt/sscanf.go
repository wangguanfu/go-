package main

import "fmt"

/*

	原型: int sscanf (const char *str,const char * format,........);

	说明： sscanf()会将参数str的字符串根据参数format字符串来转换并格式化数据。转换后的结果存于对应的参数内。

	成功则返回参数数目，失败则返回0。

	注意：sscanf与scanf类似，都是用于输入的，只是后者以键盘(stdin)为输入源，前者以固定字符串为输入源。

	大家都知道sscanf是一个很好用的函数，利用它可以从字符串中取出整数、浮点数和字符串等等。它的使用方法简单，特别对于整数和浮点数来说。

	这里就举几个经常用到的例子来说明他的用法，便于大家深刻理解他的用法

*/

func main() {

	// 对于空格分割的字符串 进行格式化赋值
	var t1 = "xmge/27/男 212 324"
	var name,sex string
	var age uint
	fmt.Println(fmt.Sscanf(t1,"%s/%d/%s",&name,&age,&sex))
	fmt.Println(name,age,sex)


}