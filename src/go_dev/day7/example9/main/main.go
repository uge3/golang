package main

import (
	"fmt"
	"os"
	"time"
)

//自定义错误接口 示例
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

//实现接口的方法
func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s op=%s createTime=%s message=%s", p.path, p.op, p.createTime, p.message)
}

//文件打开错误
func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}

func main() {
	err := Open("test.sls")
	if err != nil {
		fmt.Println(err)
	}
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}
}
