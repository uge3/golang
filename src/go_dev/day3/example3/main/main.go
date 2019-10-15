package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	var input int
	n = rand.Intn(100) //随机生成一个整数

	for {
		fmt.Println("请输入整数:")
		fmt.Scanf("%d", &input) //
		flag := false
		switch {
		case input == n:
			fmt.Println("您猜对了!")
			flag = true
		case input > n:
			fmt.Println("数字大了!")
		case input < n:
			fmt.Println("太小了!")
		}
		if flag {
			break
		}
	}
}
