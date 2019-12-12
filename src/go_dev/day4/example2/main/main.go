package main

import "fmt"

func test() {
	s1 := new([]int) //new 返回的是指针
	fmt.Println(s1)
	s2 := make([]int, 10) //make 返回的是地址
	fmt.Println(s2)
	*s1 = make([]int, 3)
	(*s1)[1] = 444
	fmt.Println(*s1)
	return
}

func main() {

	test()

}
