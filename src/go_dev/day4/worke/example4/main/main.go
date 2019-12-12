package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//大数相加 模拟
func multi(str1, str2 string) (result string) {
	//判断 两个字符串都是空的,表示就是0
	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}
	//定义 标识 用于从后往前加
	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left int
	for index1 >= 0 && index2 >= 0 {
		c1 := str1[index1] - '0' //ASII码对应的数字
		c2 := str2[index2] - '0'
		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1 //数字超过9,进位
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0' //相加后的字符串
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}
	//字符长度不一样
	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0' //相加后的字符串
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}
	for index2 >= 0 {
		c2 := str2[index2] - '0'
		sum := int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0' //相加后的字符串
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}
	if left == 1 {
		result = fmt.Sprintf("1%s", result) //如果还有进位,前面加1
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
	strSlice := strings.Split(string(result), "+") //分割字符串,以 + 号对分
	if len(strSlice) != 2 {
		fmt.Println("输入有误,请按 a + b 的格式输入")
		return
	}
	strNumber1 := strings.TrimSpace(strSlice[0]) //分别取出字符串
	strNumber2 := strings.TrimSpace(strSlice[1])
	fmt.Println(multi(strNumber1, strNumber2))
}
