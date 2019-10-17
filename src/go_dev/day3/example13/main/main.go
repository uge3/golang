package main

import "fmt"

func test(a, b int) int {
	//匿名函数
	result := func(a1, b1 int) int {
		return a1 + b1
	}(a, b)
	return result
}

func main() {
	s := test(7, 90)
	fmt.Println(s)
}
