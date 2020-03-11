package main

import (
	"bufio"
	"fmt"
	"os"
)

//带缓存区的读写
func main() {
	fmt.Println("请输入:")
	reader := bufio.NewReader(os.Stdin)
	//终端输入
	str, err := reader.ReadString('\n')
	fmt.Println("=====================")
	if err != nil {
		fmt.Println("read string failde,err:", err)
		return
	}
	fmt.Printf("read str succ,ret:%s\n", str)
}
