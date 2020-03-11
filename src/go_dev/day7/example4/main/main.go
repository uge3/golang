package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//计数结构体
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

//带缓存区的读写
func main() {

	//文件读取
	file, err := os.Open("E:\\learn\\golang\\src\\go_dev\\day7\\example\\main\\test.log")
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}
	defer file.Close() //关闭文件
	var count CharCount
	readerfile := bufio.NewReader(file)
	for {
		str, err := readerfile.ReadString('\n')
		if err == io.EOF { // io.EOF表示文件结尾
			break
		}
		if err != nil {
			fmt.Printf("read filefailed,err:%v\n", readerfile)
			break
		}
		runeArr := []rune(str)
		for _, v := range runeArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("chan count:%d\n", count.ChCount)
	fmt.Printf("num count:%d\n", count.NumCount)
	fmt.Printf("span count:%d\n", count.SpaceCount)
	fmt.Printf("othe count:%d\n", count.OtherCount)

}
