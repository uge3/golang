package main

import "fmt"

func main() {
	str := "hello,word,中国,我爱你！"
	for index, val := range str {
		fmt.Printf("下标[%d],值[%c],长度[%d] \n", index, val, len([]byte(string(val))))
	}
	fmt.Println("==================")
}
