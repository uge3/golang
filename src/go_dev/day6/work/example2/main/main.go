package main

import (
	"fmt"
)

type Books struct {
	Name string
	Age  int
}

func test(l []*Books) {
	fmt.Println(l)
}

func main() {
	var ListBook []*Books
	test(ListBook)
}
