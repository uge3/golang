package main

import (
	"fmt"
	"time"
)

type Cat struct {
	Name string `json:"name"`
	age  int
}

type Cat2 struct {
	Name string `json:"name"`
	age  int
}

//结构体可以有匿名,但一种数据类型只有一个
type Train struct {
	Cat
	start time.Time
	int
}

//结构体内部有多个结构体,可能造成冲突,赋值时不能省略中间的结构体
type Train2 struct {
	Cat
	Cat2
}

func main() {
	var t Train
	t.Name = "ttt"
	t.age = 90
	t.int = 0
	t.start = time.Now()
	fmt.Println(t)

	var t2 Train2
	t2.Cat.Name = "cat1"
	t2.Cat2.Name = "cat2"
	fmt.Println(t2)
}
