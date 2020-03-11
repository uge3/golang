package main

import "fmt"

type student struct {
	Name  string
	Age   int
	Score float32
}

//字符串格式化输入
func main() {
	var str = "name 200 675.45"
	var stu student
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)

}
