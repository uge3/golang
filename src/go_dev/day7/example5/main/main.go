package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "E:/learn/golang/src/go_dev/day7/example5/main/products.txt"
	outputFile := "E:/learn/golang/src/go_dev/day7/example5/main/products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile) //整个文件内容取入字符
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s", err)
		return
	}
	fmt.Printf("%s\n", string(buf))               //转为字符串
	err = ioutil.WriteFile(outputFile, buf, 0x64) //写入新的文件
	if err != nil {
		panic(err.Error())
	}
}
