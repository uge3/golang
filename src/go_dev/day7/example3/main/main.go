package main

import (
	"bufio"
	"fmt"
	"os"
)

//带缓存区的读写
func main() {

	//文件读取
	file, err := os.Open("E:\\learn\\golang\\src\\go_dev\\day7\\example\\main\\test.log")
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}
	defer file.Close() //关闭文件

	readerfile := bufio.NewReader(file)
	fmt.Printf("read str succ,ret:%s\n", readerfile)
	//终端输入
	fmt.Println("请输入:")
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	fmt.Println("=====================")
	if err != nil {
		fmt.Println("read string failde,err:", err)
		return
	}
	fmt.Printf("read str succ,ret:%s\n", str)
}
