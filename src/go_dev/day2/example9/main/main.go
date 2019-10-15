package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	return
}

func swap2(a int, b int) (int, int) {
	return b, a
}

func main() {
	var a = 3
	var b = 4
	var (
		c int
		d int
	)
	fmt.Println("交换前a=", a, "b=", b)
	// swap(&a, &b)
	// fmt.Println("交换后a=", a, "b=", b)
	c, d = swap2(a, b)
	fmt.Println("交换后c=", c, "d=", d)
}
