package main

import "fmt"

//空接口 可以实现任何变量
type Test interface {
}

func main() {
	var a interface{}
	var b int
	a = b
	fmt.Printf("type is a", a)
	var c float32
	a = c
	fmt.Printf("type is a", a)
}
