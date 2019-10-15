package main

import(
	"fmt"
	//"go_dev/day2/exampe2/add"
	add "go_dev/day2/exampe2/add"  //使用别名
)

func main()  {
	//add.init()
	fmt.Print("Name:",add.Name,"\n")
	fmt.Print("Age:",add.Age)
}