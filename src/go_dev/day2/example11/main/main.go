package main

import "fmt"

func test() {

	var a int8 = 127
	var b int16 = int16(a)
	fmt.Println(b)
}

func main() {

	var a int
	var b int32
	a = 10
	b = 23
	fmt.Println(a, b)
	//b = a + a  error 类型不同无法赋值
	b = b + 12
	fmt.Println(a, b)
	test()

}
