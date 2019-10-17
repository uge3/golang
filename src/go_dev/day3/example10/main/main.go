package main

import "fmt"

func add(a, b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = sum / 2
	return
}

func main() {
	a := 8
	b := 10
	c := add(a, b)
	d, e := calc(a, b)
	f, _ := calc(a, b)
	fmt.Println(c)
	fmt.Println(d, e, f)
}
