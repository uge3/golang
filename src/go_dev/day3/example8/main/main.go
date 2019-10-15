package main

import "fmt"

type add_func func(int, int) int //自定义函数类型

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// func operator(op add_func, a int, b int) int {
// 	return op(a, b)
// }
//等同于下方
func operator(op func(int, int) int, a int, b int) int {
	return op(a, b)
}
func main() {
	c := add
	fmt.Println(c)
	sum1 := c(10, 20)
	fmt.Println(sum1)
	sum2 := operator(c, 100, 200)
	sum3 := operator(sub, 200, 500)
	fmt.Println(sum2, "\n", sum3)
}
