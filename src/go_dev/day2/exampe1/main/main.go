package main

import(
	"fmt"
)

func list(n int)  {
	for index := 0; index <=n; index++ {
		fmt.Printf("%d+%d=%d\n",index, n - index, n)
	}
}
func main(){
	list(10)
}