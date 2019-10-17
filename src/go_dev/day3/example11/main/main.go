package main

import "fmt"

func add(a int, arg ...int) (c int) {
	var sum int = 0
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	c = sum + a
	return
}

func concat(a string, arg ...string) (result string) {
	result = a
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return
}

func main() {
	a := 8
	b := 10
	c := add(a, b, 12, 8, 65)
	res := concat("第一个", "  ", "--结尾!")
	fmt.Println(c)
	fmt.Println(res)
}
