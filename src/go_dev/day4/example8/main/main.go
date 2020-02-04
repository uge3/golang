package main

import "fmt"

func testSlice() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:] //切片 从下标1到最后
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	fmt.Println("原切片a:", a)
	fmt.Printf("len[%d],cap[%d]\n", len(s), cap(s))
	s = append(s, 8) //进行填充扩容
	s = append(s, 18)
	s = append(s, 28)
	fmt.Printf("长度[%d],容量[%d]\n", len(s), cap(s))
	s = append(s, 89)
	s[1] = 1000 //进行下标修改
	fmt.Println("新切片a:", a)
	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	for i, v := range s {
		fmt.Printf("下标:%p,值:%p\n", i, v)
	}
}

func testCopy() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7}
	b := make([]int, 1)
	copy(b, a) //拷贝,不会扩容
	fmt.Println(b)
}


func testByte() {
	s := "中,hello world"
	s1 := []byte(s)
	s1[3] = 222
	s1[4] = 200
	s1[5] = 45
	s1[6] = 66
	s1[7] = 97
	str := string(s1)
	fmt.Println(str)
}
func main() {
	//testSlice()
	//testCopy()
	//testString()
	testByte()
}
