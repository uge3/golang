package main

import (
	"fmt"
	"time"
)

//接口定义 继承结构体的方法
type Test interface {
	Print()
	Sleeps()
}

//结构体1
type Car struct {
	Name   string
	weight int
	higth  int
}

//结构体1方法-1
func (p Car) Print() {
	fmt.Println("name", p.Name)
	fmt.Println("weight", p.weight)
	fmt.Println("hight", p.higth)
}

//结构体1方法-2
func (p Car) Sleeps() {
	fmt.Println("Car-sleep", time.Now())
}

//结构体2
type Biek struct {
	Lin string
}

//结体体2方法-1
func (p Biek) Print() {
	fmt.Println("铃声:", p.Lin)
}

//结构体2方法-2
func (p Biek) Sleeps() {
	fmt.Println("Biek-sleep", time.Now())
}

func main() {
	var t Test
	var cars Car
	cars.Name = "五"
	cars.higth = 85
	cars.weight = 90
	fmt.Println("============")
	fmt.Println(cars)
	t = cars
	t.Print()
	t.Sleeps()
	fmt.Println("============")
	var bieks Biek = Biek{
		Lin: "铃铃.......",
	}
	t = bieks
	t.Print()
	t.Sleeps()

}
