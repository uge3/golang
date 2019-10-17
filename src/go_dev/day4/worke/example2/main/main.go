package main

import "fmt"

//回文
func palindrome(str string) bool {
	for i := 0; i < len(str); i++ { //无法对中文生效
		if i == len(str)/2 {
			break
		}
		last := len(str) - i - 1 //对应的字符串尾位置
		fmt.Printf("%c----%c", str[i], str[last])
		if str[i] != str[last] {
			return false
		}

	}
	return true
}

//可对中文生效
func palindromes(str string) bool {
	t := []rune(str) //表示字符
	length := len(t)
	for i, v := range t {
		fmt.Printf("%v--%v--%d\n", i, v, len(string(v)))
		if i == length/2 {
			break
		}
		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}
	return true
}

//反转字符串 可中文
func reverses(str string) string {
	var result string
	t := []rune(str) //按字符切片
	length := len(t)
	for i, _ := range t {
		result = result + string(t[length-i-1])
	}
	return result
}

func main() {
	var str string
	fmt.Println("请输入字符串:")
	fmt.Scanf("%sd", &str)
	result := reverses(str)
	fmt.Println("反转后:", result)
	if palindromes(str) || (result == str) {
		fmt.Println(str, "是回文")
	} else {
		fmt.Println(str, "不是回文")
	}

}
