package main

import (
	"encoding/json"
	"fmt"
)

//json tag json打包时取json:之后对应的值
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	var stu Student = Student{
		Name:  "st01",
		Age:   18,
		Score: 80,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json stu,err:", err)
		return
	}
	fmt.Println(string(data))
}
