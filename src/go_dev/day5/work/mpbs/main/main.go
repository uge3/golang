package main

import "fmt"

//冒泡排序
//       传入切片
func mpbs(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ { //从第二个元素开始，到之前循环位置结束
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j] //如果前面的元素大则进行位置对换
			}
		}
	}
}

func main() {
	a := [...]int{43, 9, 4, 90, 233, 87}
	fmt.Println(a)
	mpbs(a[:])
	fmt.Println(a)
}
