package main

import "fmt"

//选择排序
//       传入切片
func ssort(a []int) {
	for i := 0; i < len(a); i++ {
		var min = i                       //定义一个临时下标
		for j := i + 1; j < len(a); j++ { //从第二个元素开始，到位置结束
			if a[j] < a[min] {
				min = j //如果前面的元素大则进行位置对换
			}
		}
		a[i], a[min] = a[min], a[i] //进行元素对换
	}
}

func main() {
	a := [...]int{43, 9, 4, 90, 233, 87}
	fmt.Println(a)
	ssort(a[:])
	fmt.Println(a)
}
