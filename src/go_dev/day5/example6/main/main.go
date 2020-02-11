package main

import "fmt"

type Stuedut struct {
	Name  string
	Age   int
	Score float32
}

//结构体方法
func (p *Stuedut) init(name string, age int, score float32) {
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(p)
}

//可以有返回值
func (self Stuedut) get() Stuedut {
	return self
}

func main() {
	fmt.Println("============")
	var stu Stuedut

	//(&stu).init("stu1", 18, 98.9)
	stu.init("stu1", 18, 98.9)//同上
	stu1 := stu.get()
	fmt.Println(stu1)
}
