package main

import "fmt"

//菲波那契数列   非递归方式
func fba(n int) {
	var a []int64
	a = make([]int64, n)
	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
		fmt.Println(a[i])
	}
}

func main() {
	fba(50)
}
