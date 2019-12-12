package main

import "fmt"

func testArray() {
	var age0 [5]int = [5]int{1, 2, 3}
	var age1 = [5]int{1, 2, 3, 4, 5}
	var age2 = [...]int{1, 2, 3, 4, 5, 6}
	var str = [5]string{3: "hello world", 4: "tom"}
	fmt.Println(age0)
	fmt.Println(age1)
	fmt.Println(age2)
	fmt.Println(str)
}
func testArrays() {
	//多维数组
	var age [2][3]int
	var age3 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(age)
	fmt.Println(age3)
	//多维数组的遍历
	for row, v := range age3 {
		for col, v1 := range v {
			fmt.Printf("%d,%d=%d ", row, col, v1)
		}
		fmt.Println()
	}
}

func main() {
	testArray()
	testArrays()
}
