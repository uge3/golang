package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name  string
	Title string
	age   string
}

func main() {
	t, err := template.ParseFiles("E:\\learn\\golang\\src\\go_dev\\day10\\http_template\\index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "may", age: "31", Title: "我的网端"}
	//文件 写入终端 os.Stdout
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
