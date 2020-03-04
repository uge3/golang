package main

import (
	"fmt"
	"reflect"
)

//结构体
type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

//结构体方法
func (s Student) Set(name string, age int, score float32) {
	s.Name = name
	s.Age = age
	s.Score = score
}

//反射获取方法
func TestStruct(b interface{}) {
	val := reflect.ValueOf(b) //获取值
	kd := val.Kind()
	if kd != reflect.Struct { //如果不是对应的类型,结构体
		fmt.Println("expect,struct")
		return
	}
	num := val.NumField() //获取字段的数量
	for i:=0;i<num;i++{
		fmt.Printf("%d,%v\n",i,val.Field(i).Kind())//输出字段的值类型
	}
	fmt.Printf("结构体字段 %d\n", num)
	numOfMethod := val.NumMethod() //获取结构体方法
	fmt.Printf("结构体方法 %d\n", numOfMethod)
	var params []reflect.Value
	val.Method(0).Call(params)

}

func main() {
	fmt.Println()
	var a Student = Student{
		Name:  "stu001",
		Age:   18,
		Score: 98,
	}
	TestStruct(a)
	fmt.Println(a)

}
