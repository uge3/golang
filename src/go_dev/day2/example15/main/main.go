package main

import "fmt"

//反转字符串
func reverse(str string) string {
	var result string
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		result = result + fmt.Sprintf("%c", str[strlen-i-1]) //切片 从后往前取字符
	}
	return result
}

func main() {
	str1 := "hello"
	str2 := "world"
	str3 := str1 + " " + str2
	str4 := fmt.Sprintf("%s %S", str1, str2) //格式化拼接
	n := len(str3)                           //长度
	substr := str4[0:4]                      //切片
	fmt.Println("str3 len:", n)
	fmt.Println(str3)
	fmt.Println(substr)
	str5 := reverse(str3)
	fmt.Println(str5)
}
