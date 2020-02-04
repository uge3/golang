package main

import "fmt"

//插入排序
//       传入切片
func mpbs(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0; j-- { //从之前循环位置开始，往最前面
			if a[j] > a[j-1] { //如果后面的元素大则不变
				break
			}
			a[j], a[j-1] = a[j-1], a[j] //如果前面的元素大则进行位置对换
		}
	}
}

func main() {
	a := [...]int{43, 9, 4, 90, 233, 87}
	fmt.Println(a)
	mpbs(a[:])
	fmt.Println(a)
}
