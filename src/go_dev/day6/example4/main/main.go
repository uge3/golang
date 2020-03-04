package main

import "fmt"

//接口一  读
type Reader interface {
	Read()
}

//接口二 写
type Writer interface {
	Writer()
}

//接口三 读写
type ReaderWriter interface {
	Reader
	Writer
}

// 文件结构体
type File struct {
}

// 对应方法
func (p *File) Read() {
	fmt.Println("file,read data")
}

// 对应方法
func (p *File) Writer() {
	fmt.Println("file,Writer data")
}
func Test(rw ReaderWriter) {
	rw.Read()
	rw.Writer()
}

func main() {
	var f File
	// Test(&f)
	var b interface{}
	b = f
	v, ok := b.(ReaderWriter)
	fmt.Println("____", v, ok)

}
