package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//结构体
type Student struct {
	Name  string `json:"student_name"`
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
	tye := reflect.TypeOf(b)  //获取类型
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct { //如果不是指针并且是对应的类型,结构体
		fmt.Println("expect,struct")
		return
	}
	num := val.Elem().NumField() //获取字段的数量
	for i := 0; i < num; i++ {
		fmt.Printf("%d,%v\n", i, val.Elem().Field(i).Kind()) //输出字段的值类型
	}
	fmt.Printf("结构体字段 %d\n", num)

	val.Elem().Field(1).SetInt(33) //通过反射进行修改

	tag := tye.Elem().Field(0).Tag.Get("json") //获取json格式的字段
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := val.Elem().NumMethod() //获取结构体方法
	fmt.Printf("结构体方法 %d\n", numOfMethod)
	var params []reflect.Value
	val.Elem().Method(0).Call(params)

}

func main() {
	fmt.Println()
	var a Student = Student{
		Name:  "stu001",
		Age:   18,
		Score: 98,
	}

	result, _ := json.Marshal(a) //json数据格式
	fmt.Println("json result:", string(result))
	TestStruct(&a)
	fmt.Println(a)

}
