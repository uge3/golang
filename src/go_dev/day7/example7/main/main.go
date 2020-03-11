package main

import (
	"flag"
	"fmt"
	"os"
)

//命令行参数：是string切片 ---打印解析命令行参数
func main() {
	fmt.Print("len of args:%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}

	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please,input conf path")
	flag.IntVar(&logLevel, "d", 0, "please input log level")
	flag.Parse() //生效
	fmt.Println("path:",confPath)
	fmt.Println("log level:",logLevel)
}
