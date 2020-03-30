package main

//线程 goroutine  调度模型
import (
	"fmt"
	"runtime"
	"sync" //锁
	"time"
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
	fmt.Println(t.n, sum)
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	//启用多线程
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	for i := 0; i < 25; i++ {
		t := &task{n: i}
		go cale(t)
	}
	time.Sleep(time.Second * 5)

}
