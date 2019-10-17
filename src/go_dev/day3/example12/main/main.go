package main

import "fmt"

func add(a int) int {
	var b int
	defer fmt.Println(a)
	a = 90
	b = a + 1
	fmt.Println(b)
	return b
}
// func read() {
// 	file := open(filename)
// 	defer file.Close() //文件句柄关闭
// 	mc.Lock()
// 	defer mc.Unlock() //锁释放
// 	conn := openDatabase()
// 	defer conn.Close() //数据库连接释放
// }

func main() {
	s := add(7)
	fmt.Println(s)
}
