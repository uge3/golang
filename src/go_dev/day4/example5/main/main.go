package main

import "fmt"

//数组
func test1() {
	var a [10]int
	a[0] = 1000
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])

	}
	for index, val := range a {
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}

//数组是值类型，因此改变副本的值，不会改变本身的值
func test2() {
	var a [5]int
	b := a
	b[0] = 56
	fmt.Println("数组A:", a[0])
	fmt.Println("数组B:", b[0])
}

//需要改变本身的值,传入指针
func test3(arr *[6]int) {
	arr[1] = 345
}

func main() {
	test1()
	test2()
	var a [6]int
	a[1] = 999
	fmt.Println("原赋值:", a[1])
	test3(&a)
	fmt.Println("修改后:", a[1])

}
