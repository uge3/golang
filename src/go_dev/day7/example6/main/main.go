package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

//压缩文件的读写
func main() {
	fName := "E:/learn/golang/src/go_dev/day7/example6/main/products.txt.gz"
	var r *bufio.Reader //创建一个文件读写句柄
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v,Can't open %s:error :%s\n", os.Args[0], fName, err)
		os.Exit(1) //如果出错文件句柄关闭
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open gzip failed,err;%v\n", err)
		return
	}
	r = bufio.NewReader(fz) //将解压用的文件放入文件柄进行操作
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
	defer fi.Close()
}
