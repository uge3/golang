package main

import "fmt"

//接口
type Carer interface {
	GetName() string
	Run()
	DiDi()
}

//接口2
type Test interface {
	Hell()
}

//结构体  1 实现接口,需要实现接口内的所有方法
type BMW struct {
	Name string
}

// 方法一
func (p *BMW) GetName() string {
	return p.Name
}

//方法二
func (p *BMW) Run() {
	fmt.Printf("car name:%s,RUN", p.Name)

}

//方法三
func (p *BMW) DiDi() {
	fmt.Println("滴滴")
}

//方法四
func (p *BMW) Hell() {
	fmt.Printf("this is %s", p.Name)
}

//结构体  2 实现接口,需要实现接口内的所有方法
type BYD struct {
	Name string
}

// 方法一
func (p *BYD) GetName() string {
	return p.Name
}

//方法二
func (p *BYD) Run() {
	fmt.Printf("car name:%s,RUN", p.Name)
}

//方法三
func (p *BYD) DiDi() {
	fmt.Println("滴滴----")
}

func main() {
	var car Carer //设置 接口
	var test Test
	fmt.Println(car)
	bmw := &BMW{
		Name: "bmw",
	}
	car = bmw
	car.GetName()
	car.Run()
	car.DiDi()
	test = bmw
	test.Hell()
	var byd BYD
	byd.Name = "BYD"
	car = &byd
	car.Run()
	car.DiDi()
}
