package main

import "fmt"

func main() {
	str := "test\ntest"
	str2 := `
	床前明用光，\n
	地上鞋三双。
	`
	fmt.Println(str)
	fmt.Println(str2)
}
