package main

import (
	"fmt"
	link "go_dev/day6/example6/link"
)

func main() {
	// 声明一个链表
	var intLink link.Link
	for i := 0; i < 10; i++ {
		intLink.IntsertHead(i) //头部插入
	}
	intLink.Tarns()

	fmt.Println("=================")
	for i := 1; i < 20; i++ {
		intLink.IntsertTail(fmt.Sprintf("str %d", i)) //尾部插入
	}
	intLink.Tarns()
}
