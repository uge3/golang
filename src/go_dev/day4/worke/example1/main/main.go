//编写程序  在终端输出九九乘法表.
package main

import "fmt"

func multi() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

//一个数如果恰 好等 于它的因子之和,这个数就称为"完数",例如:6/1=6,6/2=3,6/3=2, 6=1+2+3,
func perfect(n int) bool {
	var sum int = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return n == sum
}
func perfectss(n int) {
	for i := 1; i < n+1; i++ {
		if perfect(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	multi()
	var n int
	fmt.Println("请输入整数:")
	fmt.Scanf("%d", &n)
	perfectss(n)

}
