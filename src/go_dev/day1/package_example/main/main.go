package main



import (
	"go_dev/day1/package_example/calc"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go calc.Test_goroute(i)
	}
	time.Sleep(time.Second)
	sum := calc.Add(100,300)
	fmt.Println("sum=",sum)
}
