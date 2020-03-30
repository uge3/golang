package main

//线程 goroutine  调度模型
import (
	"fmt"
	"time"
)

func test() {
	var i int
	for {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}
}

func main() {
	go test() //启动一个线程  test()为协程
	//time.Sleep(time.Second * 10)
	for {
		fmt.Println("i,running in main")
		time.Sleep(time.Second)
	}

}
