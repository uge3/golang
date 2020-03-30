package main

//线程 goroutine  调度模型
import (
	"fmt"
	"sync" //锁
)

//全局变量
var (
	m    = make(map[int]uint64)
	lock sync.Mutex //写多用互拆锁
)

type task struct {
	n int
}

func cale(t *task) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i := 0; i < 25; i++ {
		t := &task{n: i}
		go cale(t)
		fmt.Printf("写入内容：%\n", i)
	}
	//time.Sleep(10 * time.Second)
	lock.Lock()
	for k, v := range m {
		fmt.Printf("%d != %v\n", k, v)
	}
	lock.Unlock()
}
