package main

import (
	"fmt"
)

func modlfr(a int) { //传入值，为副本
	a = 10
	return
}
func modlfr1(a *int) { //传入地址
	*a = 10 //修改地址指向的内存进行修改
}

func main() {
	var a = 100
	var b chan int = make(chan int, 1) // b:= make(chan int ,1)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	modlfr(a)
	fmt.Println("a=", a)
	modlfr1(&a)
	fmt.Println("&a=", a)

}
