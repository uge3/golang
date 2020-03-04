package main

import "fmt"

type Stu struct {
	Name string
	Age  string
}

func Test(a interface{}) {
	b := a.(int) //接口转类型
	b += 3
	fmt.Println(b)
}

func TestStu(s interface{}) {
	c, ok := s.(Stu) //ok 用于接收错误!
	if ok == false {
		fmt.Println("错误!")
		return
	}
	fmt.Println(c)
}

//多类型判断
func Just(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%d个是,bool类型,值为%v\n", index, v)
		case int:
			fmt.Printf("第%d个是,int类型,值为%v\n", index, v)
		case string:
			fmt.Printf("第%d个是,string类型,值为%v\n", index, v)
		case Stu:
			fmt.Printf("第%d个是,struct类型,值为%v\n", index, v)
		}
	}
}

func main() {
	var a int
	Test(a)
	var s Stu
	TestStu(s)
	TestStu(a)
	Just(a, "ss", 234, s)

}
