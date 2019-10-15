package main

import (
	"fmt"
	"strconv"
	"strings"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func main() {
	var (
		url  string
		path string
	)
	fmt.Println("请输入url , path")
	fmt.Scanf("%s%s", &url, &path) //终端输入
	url = urlProcess(url)
	path = pathProcess(path)
	fmt.Println("url:", url, "\n", "path:", path)
	str := strings.Repeat("str", 4)       //字符串重复次数
	str1 := strings.Fields("abc efg hij") //返回以空格分隔的所有子字符串
	str2 := strings.Join(str1, "--")      //用指定字符把字符串列表中的元素拼接起来

	str3 := strconv.Itoa(100)        //把一个整数 转为字符串
	nums, err := strconv.Atoi("888") //把 字符串转为一个整数
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(nums, err)
}
