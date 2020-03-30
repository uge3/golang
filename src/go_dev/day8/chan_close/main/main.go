package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	for {
		var b int
		b, ok := <-ch
		//用于判断管道关闭
		if ok == false {
			fmt.Println("chan is nil")
			break
		}
		fmt.Println(b)
		time.Sleep(time.Second)
	}
}
