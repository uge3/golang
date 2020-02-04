package main

import "fmt"

//结构体
type Stuedt struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var stu Stuedt
	stu.Name = "颜靖靖"
	stu.Age = 39
	stu.Score = 99.87
	fmt.Println(stu)
	fmt.Printf("Name:%p\n", &stu.Name)
	fmt.Printf("Name:%p\n", &stu.Age)
	fmt.Printf("Name:%p\n", &stu.Score)
	//初始化方法二
	var stu2 *Stuedt = &Stuedt{
		Age:  40,
		Name: "大规模",
	}
	fmt.Println(stu2.Name)
	//初始化方法三
	var stu3 = Stuedt{
		Age:  30,
		Name: "大规模2",
	}
	fmt.Println(stu3.Name)

}
