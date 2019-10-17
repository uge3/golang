package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(str string) (wordCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			// wordCount++
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}
func main() {
	// var str string
	fmt.Println("请输入字符串:")
	// fmt.Scanf("%sd", &str)
	reader := bufio.NewReader(os.Stdin) //系统标准终端
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}
	wordCount, spaceCount, numberCount, otherCount := count(string(result))
	fmt.Printf("英文字符:%d,空格:%d,数字:%d,其它:%d.", wordCount, spaceCount, numberCount, otherCount)

}
