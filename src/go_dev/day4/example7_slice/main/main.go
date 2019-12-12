package main

import "fmt"

func testSlice() {
	var slice []int
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	slice = arr[2:5] //切片，包括开始下标的元素，不包括结束下标的元素
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = slice[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = arr[:4]
	fmt.Println(slice)
}

func main() {
	testSlice()

}
