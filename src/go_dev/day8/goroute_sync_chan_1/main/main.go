package main

//chan 管道同步
import (
	"fmt"
)

func calc(taskChan chan int, resChan chan int, exitChan chan bool) {
	for v := range taskChan {
		fmt.Println("当前检验数字：", v)
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resChan <- v
		}
	}
	fmt.Println("exit,退出")
	exitChan <- true //协程结果放出结果管道

}

func main() {
	intChan := make(chan int, 1000) //定义一个管道
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	go func() {
		for i := 0; i < 100000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	//启用8个协程
	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	//等待所有计算的goroutine全部退出
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
			fmt.Println("协程结束！", i)
		}
		close(resultChan)
	}()

	//使用匿名函数
	//go func() {
		for v := range resultChan {
			fmt.Println("素数：", v)
		}
	 //}()
}
