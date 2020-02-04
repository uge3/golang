package main

import (
	"fmt"
	"sort"
)

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[0] = 10
	a[1] = 101
	a[2] = 8
	a[3] = 77
	a[4] = 1018
	var keys []int
	for k, v := range a {
		fmt.Println(k, v, "\n")
	}
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys[:]) //进行排序
	fmt.Println("----------------我是分割线---------------------")
	for _, v := range keys {
		fmt.Println(v, a[v], "\n")
	}
}

//map key vula 反转
func testMapSort() {
	a := make(map[string]int)
	b := make(map[int]string)
	a["a"] = 101
	a["b"] = 98
	a["c"] = 77
	for k, v := range a {
		b[v] = k
	}
	fmt.Println(b)
}

func main() {
	// testMap()
	testMapSort()
}
