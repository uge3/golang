package main

import (
	"errors"
	"fmt"
	"time"
)

func initConfig() (err error) {
	return errors.New("初始化错误")
}

func test() {
	//匿名函数  panic 和 recover:用来做错误处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	b := 0
	a := 100 / b
	fmt.Println(a)
	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}

func main() {

	test()
	time.Sleep(time.Second)

	var i int
	fmt.Println(i)
	j := new(int) //new 返回的是指针
	*j = 900
	fmt.Println(*j)

	var a []int
	a = append(a, 10, 19, 45, 98)
	a = append(a, a...) //... 可以当可变参数,以及展开
	fmt.Println(a)

}
