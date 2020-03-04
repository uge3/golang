package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b) //进行反射应用
	fmt.Println(t)

	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)
	iv := v.Interface()     //转成接口
	stu, ok := iv.(Student) //进行断言
	if ok {
		fmt.Printf("%v,%T\n", stu, stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b) //获取值
	val.Elem().SetInt(100)    //elem()通过地址取值
	c := val.Elem().Int()
	fmt.Printf("获取值,接口,%d\n", c)
}

func main() {
	fmt.Println()
	var a Student = Student{
		Name:  "stu001",
		Age:   18,
		Score: 98,
	}
	test(a)
	var b int = 1
	testInt(&b)
	fmt.Println(b)
}
