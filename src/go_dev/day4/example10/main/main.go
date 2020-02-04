package main

import "fmt"

func testMap() {
	var a map[int]int
	a = make(map[int]int, 1) //make初始化
	//a:=make(map[string]sting , 1)  //效果同上
	a[1] = 1
	a[2] = 10
	a[3] = 100
	fmt.Println(a)
}

func testMaps() {
	a := make(map[string]map[string]map[string]string, 3)

	a["第一层"] = make(map[string]map[string]string)
	a["第一层"]["第二层"] = make(map[string]string)
	a["第一层"]["第二层"]["第三层"] = "第四层"
	a["第一层"]["第二层"]["第三层2"] = "第四层2"
	//新二层初始化
	a["第一层"]["第二层2"] = make(map[string]string)
	a["第一层"]["第二层2"]["第三层"] = "第四层"
	//新一层初始化
	a["第一层2"] = make(map[string]map[string]string)
	a["第一层2"]["2第二层"] = make(map[string]string)
	a["第一层2"]["2第二层"]["2第三层"] = "2第四层"
	fmt.Println(a)
	modify2(a)
	fmt.Println("新增后:", a)
}

func modify(a map[string]map[string]string) {
	_, ok := a["第一层"]
	if !ok {
		a["第一层"] = make(map[string]string)
	}
	a["第一层"]["password"] = "123456"
	a["第一层"]["nickname"] = "昵称1"
	return
}

func modify2(a map[string]map[string]map[string]string) {
	valu, ok := a["一"]
	if ok {
		_, ok2 := valu["二"]
		if !ok2 {
			modify(valu)
		}

	} else {
		a["一"] = make(map[string]map[string]string)
		a["一"]["二层"] = make(map[string]string)
		modify(a["一"])
	}
	return
}

func testMap2() {
	a := make(map[string]map[string]string, 3)
	modify(a)
	fmt.Println(a)
}

func trans(a map[string]map[string]string) {
	for k, v := range a {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println("\t", k1, v1)
		}
	}
}
func testMap4() {
	a := make(map[string]map[string]string, 100)
	a["1"] = make(map[string]string)
	a["1"]["第二层1"] = "1a" //插入
	a["1"]["第二层2"] = "1b"
	a["1"]["第二层3"] = "1c"
	a["2"] = make(map[string]string)
	a["2"]["2第二层1"] = "2a"
	a["2"]["2第二层2"] = "2b"
	a["2"]["2第二层3"] = "2c"
	trans(a)
	delete(a, "2") //删除
	fmt.Println("----------------我是分割线---------------------")
	trans(a)
}
func testMap5() {
	var a []map[int]int
	a = make([]map[int]int, 5)
	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][1] = 10
	fmt.Println(a)
}

func main() {
	// testMap()
	// testMaps()
	// testMap2()
	//testMap4()
	testMap5()
}
