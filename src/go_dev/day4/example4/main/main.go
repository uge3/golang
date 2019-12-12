package main

import (
	"fmt"
	"strings"
)

//闭包       返回一个函数
func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

//示例,后缀名  定义要添加的后缀名 suffix
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false { //判断是否有对应的后缀名,如果没有进行添加
			return name + suffix
		}
		return name
	}
}

func main() {
	f := Adder()        //可以运用于累计
	fmt.Println(f(1))   //内部的X 为1  保持不变
	fmt.Println(f(200)) //加上本次的200 值为 201
	fmt.Println(f(3000))
	fmt.Println("=============================")
	g := makeSuffix(".bmp")
	fmt.Println(g("test"))
	fmt.Println(g("pic"))
	fmt.Println(g("sss.bmp"))

}
