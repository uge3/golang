package main

import "fmt"

//channel 的定义
type student struct {
	name string
}

func main() {
	var stuChan chan interface{}        //定义一个管道
	stuChan = make(chan interface{}, 2) //初始化管道
	stu := student{name: "stu01"}
	stuChan <- &stu //存入管道

	var stu01 interface{}
	stu01 = <-stuChan//从管道取出
	var stu02 *student
	stu02, ok := stu01.(*student)//接口转回结构体
	if !ok {
		fmt.Println("can not convert")
		return
	}
	fmt.Println(stu02)
}
