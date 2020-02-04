package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex     //互拆锁
var rwlock sync.RWMutex //文件读写锁

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[0] = 10
	a[1] = 101
	a[2] = 8
	a[3] = 77
	a[4] = 1018
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[3] = (rand.Intn(1000))/2 + 3*(rand.Intn(89))
			lock.Unlock()
		}(a)
	}
	lock.Lock() //  互拆锁  读写的位置都需要加锁
	fmt.Println(a)
	lock.Unlock()
	time.Sleep(time.Second)
}

func testMapRwlock() {
	var a map[int]int
	var count int32
	a = make(map[int]int, 5)
	a[0] = 10
	a[1] = 101
	a[2] = 8
	a[3] = 77
	a[4] = 1018
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[3] = b[3] + (rand.Intn(1000))/2 + 3*(rand.Intn(89))
			rwlock.Unlock()
		}(a)
	}
	for i := 0; i < 20; i++ {
		go func(b map[int]int) {
			rwlock.RLock()
			fmt.Println(a)
			rwlock.RUnlock()
			atomic.AddInt32(&count, 1) //进行原子操作
			//fmt.Println(atomic.LoadInt32(&count))
		}(a)
	}

	time.Sleep(time.Second * 5)
	fmt.Println(atomic.LoadInt32(&count))
}

func testRwlockvsLock() {
	var a map[int]int
	var count int32
	a = make(map[int]int, 5)
	a[0] = 10
	a[1] = 101
	a[2] = 8
	a[3] = 77
	a[4] = 1018
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			//rwlock.Lock()
			lock.Lock()
			b[3] = rand.Intn(89)
			time.Sleep(10 * time.Millisecond)
			//rwlock.Unlock()
			lock.Unlock()
		}(a)
	}
	for i := 0; i < 20; i++ {
		go func(b map[int]int) {
			for {
				//rwlock.RLock()
				lock.Lock()
				time.Sleep(time.Millisecond)
				lock.Unlock()
				//rwlock.RUnlock()
				atomic.AddInt32(&count, 1) //进行原子操作
			}

		}(a)
	}

	time.Sleep(time.Second * 5)
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	//testMap()
	//testMapRwlock()
	testRwlockvsLock()
}
