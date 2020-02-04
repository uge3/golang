package main

import "fmt"

//快速排序
//         传入切片  下标    最后第二个元素下标
func qsort(a []int, left int, right int) {
	if left >= right { //如果当前切片只有二个或一个元素直接返回
		return
	}
	valu := a[left]                      //定义左边第一人元素值
	k := left                            //定义临时元素的下标为左边第一个下标
	for i := left + 1; i <= right; i++ { //从左边第二个起到最后一个元素进行对比
		if a[i] < valu { //如果当前元素小于第一个元素
			a[k] = a[i]   //当前元素与第一个元素对换
			a[i] = a[k+1] //然后交后的元素与第二个元素对换
			k++           //位置下标加1
		}
	}
	a[k] = valu          //确定当前第一元素所上的位置
	qsort(a, left, k-1)  //左边递归
	qsort(a, k+1, right) //右边递归
}

func main() {
	a := [...]int{43, 9, 4, 90, 233, 87}
	fmt.Println(a)
	qsort(a[:], 0, len(a)-1) //传入切片，初始下标0，长度-1
	fmt.Println(a)
}
