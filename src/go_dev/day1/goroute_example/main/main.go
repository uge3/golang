package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute"
)

func main() {
	var pipe chan int        //生成一个管道
	pipe = make(chan int, 1) //函数返回值 放入管道
	go goroute.Add(100, 300, pipe)

	sum := <-pipe
	fmt.Println("sum=", sum)
}
