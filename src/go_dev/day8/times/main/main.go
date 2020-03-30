package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 0; i < 1024; i++ {
		go func() {
			for {
				select {
				//定时器,不会被回收,不建议使用 time.After
				case <-time.After(time.Second):
					fmt.Println("i")
				}
			}
		}()
	}
	time.Sleep(time.Second * 100)

}
