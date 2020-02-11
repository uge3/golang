package main

import "fmt"

type Car struct {
	Name   string
	weight int
	Score  float32
}

//结构体方法
func (p *Car) rung() {

	fmt.Println("ruunng")
}

//继承
type Bike struct {
	Car
	Name string
}

//继承2
type Train struct {
	c Car
}

//格式化输出接口
func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s],weight=[%d]", p.c.Name, p.c.weight)
	return str
}

func main() {
	fmt.Println("============")
	var bark Bike
	bark.weight = 750
	bark.Name = "bike"
	fmt.Println(bark)
	bark.rung()

	var tan Train
	tan.c.Name = "tanm"
	fmt.Println(tan)
	tan.c.weight = 850
	tan.c.Score = 98.36

	tan.c.rung()
	fmt.Printf("%s", &tan)

}
