package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Person struct {
	Name  string
	Title string
	Age   int
}

//全局模板
var myTemplate *template.Template

//自定义终端输出
type Result struct {
	output string
}

func (p *Result) Write(b []byte) (n int, err error) {
	fmt.Println("called by template")
	p.output += string(b)
	return len(b), nil
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("userinfo 页面")
	fmt.Fprintf(w, "用户信息端！")
	p := Person{Name: "may", Age: 31, Title: "我的网端"}

	myTemplate.Execute(w, p) //模板返回到网页端
	//打开 文件 ，没有创建
	file, err := os.OpenFile("E:/testfile/day10_http_template_server.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}
	myTemplate.Execute(file, p) //模板渲染到文件
	defer file.Close()          //关闭文件

	//使用自定义终端
	resultWriter := &Result{}
	myTemplate.Execute(resultWriter, p)
	fmt.Println("template reandr data:", resultWriter.output)
}

func usersInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("用户组页面：usersInfo")
	var arr []Person
	p := Person{Name: "may1", Age: 11, Title: "我的网端1"}
	p2 := Person{Name: "may2", Age: 21, Title: "我的网端2"}
	p3 := Person{Name: "may3", Age: 31, Title: "我的网端2"}
	arr = append(arr, p)
	arr = append(arr, p2)
	arr = append(arr, p3)
	myTemplate.Execute(w, arr)
}

//模板服务启动加载函数
func initTemplat(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	//fmt.Println("页面：", myTemplate)
	return
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	fmt.Fprintf(w, "hello111") //网页端反馈
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("用户登录。。")
	fmt.Fprintf(w, "login:用户登录界面！")
}

func main() {
	fmt.Println("server start....")
	//初始化模板
	initTemplat("E:/learn/golang/src/go_dev/day10/http_template_server/index.html")
	//返回的网页端
	http.HandleFunc("/", Hello)
	http.HandleFunc("/userinfo", userInfo)
	http.HandleFunc("/usersinfo", usersInfo)
	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("0.0.0.0:8800", nil)
	if err != nil {
		fmt.Println("http listen failed!")
	}
}
