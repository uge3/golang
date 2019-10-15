package main

import "fmt"

func modify(a *int) {
	*a = 100 //通过引用地址修改其中的值
}

func main() {
	a := 8
	fmt.Println(a)
	modify(&a) //传递引用地址
	fmt.Println(a)
}
