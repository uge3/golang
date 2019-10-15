package main

import "fmt"

func main() {
	fmt.Println("============ ")
	var a int = 3
	switch a {
	case 0:
		fmt.Println("a is equal 0") //不用back
		fallthrough                 // fallthrough穿透,执行下面的代码
	case 1:
		fmt.Println("a is equal 1")
	case 2, 3, 4, 5, 6: //多条件
		fmt.Println("a is equal 2,3,4,5,6")
	case 10:
		fmt.Println("a is equal 10")
	default:
		fmt.Println("a is equal default")
	}
}
