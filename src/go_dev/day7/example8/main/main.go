package main

import (
	"encoding/json"
	"fmt"
)

//json协议 示例
type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

//结构体测试
func testStructJson() (ret string, err error) {
	user1 := &User{
		UserName: "user5",
		NickName: "呢称2",
		Age:      18,
		Birthday: "2000/1/1",
		Sex:      "男",
		Email:    "user@qq.com",
		Phone:    "00132564879",
	}
	data, err := json.Marshal(user1) //转换json格式
	if err != nil {
		fmt.Printf("json.marshal failed,err:", err)
		return
	}
	fmt.Printf("序列化、反序列化测试%s\n", string(data))
	ret = string(data)
	return
}

//结构体测试
func testStruct() {
	user1 := &User{
		UserName: "user1",
		NickName: "呢称一",
		Age:      28,
		Birthday: "2000/1/1",
		Sex:      "男",
		Email:    "user@qq.com",
		Phone:    "132564879",
	}
	data, err := json.Marshal(user1) //转换json格式
	if err != nil {
		fmt.Printf("json.marshal failed,err:", err)
		return
	}
	fmt.Printf("结构体测试%s\n", string(data))

}

//map格式
func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "usermap"
	m["age"] = 29
	m["sex"] = "man"
	data, err := json.Marshal(m) //转换json格式
	if err != nil {
		fmt.Printf("json.marshal failed,err:", err)
		return
	}
	fmt.Printf("map格式%s\n", string(data))
}

//slice 切片格式
func testSlice() {
	var s []map[string]interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "userslice1"
	m["age"] = 29
	m["sex"] = "man"
	s = append(s, m)

	m = make(map[string]interface{})
	m["username"] = "userslice2"
	m["age"] = 33
	m["sex"] = "wman"
	s = append(s, m)
	data, err := json.Marshal(s) //转换json格式
	if err != nil {
		fmt.Printf("json.marshal failed,err:", err)
		return
	}
	fmt.Printf("slice 切片格式%s\n", string(data))
}

//json协议 int格式
func testInt() {
	var age = 100
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Printf("json.marshal failed,err:", err)
		return
	}
	fmt.Printf("int格式%s\n", string(data))
}

//反序列化
func testJson() {
	data, err := testStructJson()
	if err != nil {
		fmt.Println("testJson struct failed,", err)
		return
	}
	var user2 User
	err = json.Unmarshal([]byte(data), &user2)
	if err != nil {
		fmt.Println("unmarshal failed,", err)
		return
	}
	fmt.Println(user2)
}

//json转换 序列化
func main() {
	testStruct()
	testInt()
	testMap()
	testSlice()
	fmt.Println("======================")
	testJson()
}
