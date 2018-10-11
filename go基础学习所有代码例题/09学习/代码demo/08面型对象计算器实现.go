package main

import "fmt"

//定义接口
type JIEKOU interface {
	//方法声明
	JieGuo() (int)
}

type FULEI struct {
	shuzi1 int
	shuzi2 int
}

type JIAFA struct {
	FULEI
}
type JIANFA struct {
	FULEI
}

func (jiafa *JIAFA) JieGuo() int {
	return jiafa.shuzi1 + jiafa.shuzi2
}

func (jianfa *JIANFA) JieGuo() int {
	return jianfa.shuzi1 - jianfa.shuzi2
}

type GONGCHANG struct {
}

func DuoTai(j JIEKOU) (zhi int) {
	zhi = j.JieGuo()
	return
}

func (gc *GONGCHANG) YunSuan(shuzi1 int, shuzi2 int, fuhao string) (zhi int) {

	switch fuhao {
	case "+":
		var j JIAFA = JIAFA{FULEI{shuzi1, shuzi2}}
		zhi = DuoTai(&j)
	case "-":
		var j JIANFA = JIANFA{FULEI{shuzi1, shuzi2}}
		zhi = DuoTai(&j)
	}
	return

}

func main() {

	var gc GONGCHANG

	zhi := gc.YunSuan(10, 20, "+")
	fmt.Println(zhi)
}
