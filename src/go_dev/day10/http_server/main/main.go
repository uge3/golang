package main

import (
	"fmt"
	"net/http"
)

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
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("0.0.0.0:8800", nil)
	if err != nil {
		fmt.Println("http listen failed!")
	}
}
