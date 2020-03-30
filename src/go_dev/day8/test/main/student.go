package main

import (
	"encoding/json"
	"io/ioutil"
)

type student struct {
	Name string
	Sex  string
	Age  int
}

//保存文件
func (p *student) Save() (err error) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	ioutil.WriteFile("e://testfile//day8.dat", data, 0755) //写入文件
	return
}

//打开文件
func (p *student) Load() (err error) {
	data, err := ioutil.ReadFile("e://testfile//day8.dat")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, p)
	return
}
