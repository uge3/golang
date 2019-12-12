package main

import (
	"fmt"
	"time"
)

func test(n int) int {

	fmt.Println("一直在递归", n)
	time.Sleep(time.Second)
	if n > 2 {
		return n
	}
	return test(n + 1) //调用自己
}

func cale(n int) string {
	str := "倒计时完成!"
	if n == 0 {
		return str
	}
	fmt.Println("倒计时:", n)
	time.Sleep(time.Second)
	return cale(n - 1)

}

//菲波那数
func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}
func main() {

	fmt.Println(cale(test(1)))
	for i := 0; i < 10; i++ {
		fmt.Println(fab(i))
	}
}
