package main

import "fmt"

//素数
func prime(num int) int {
	bools := 1
	for i := 2; i < num; i++ {
		if num%i == 0 {
			bools = 0
		}
	}
	return bools
}

//水仙花数
func flawes(num int) bool {
	var a, b, c int
	a = num % 10
	b = (num / 10) % 10
	c = (num / 100) % 10
	sum := a*a*a + b*b*b + c*c*c
	return sum == num
}

//阶乘之和
func sums(n int) uint64 {

	var s uint64 = 1
	var sum uint64 = 0

	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		fmt.Printf("%d!=%v\n", i, s)
		sum += s
	}
	return sum
}

func main() {
	var bools int
	//var str string
	var lens []int
	var (
		n int
		m int
	)
	fmt.Println("输入两个数: (如:100 200)")
	fmt.Scanf("%d %d", &n, &m)
	for i := n; i < m; i++ {
		bools = prime(i)
		if bools == 1 {
			//str = fmt.Sprintf("[%s] 是素数!", i)
			//temp := []byte(string(i))
			//fmt.Println(str)
			lens = append(lens, i)
			continue
		} else {
			//str = fmt.Sprintf("[%s] 不是素数!", i)
			//fmt.Println(str)
		}
	}
	fmt.Println("素数列表:", lens)
	var flawesnum []int
	for i := n; i < m; i++ {
		if flawes(i) == true {
			flawesnum = append(flawesnum, i)
		}
	}

	fmt.Println("水仙花数", flawesnum)
	s := sums(n)
	fmt.Println(s)

}
