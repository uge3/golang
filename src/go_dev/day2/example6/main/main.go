package main

import(
	"fmt"
	"time"
)

const(
	Female = 2
	Man = 1
)

func main()  {
	
	for i := 0; i < 20; i++ {
		second := time.Now().Unix()  //当前时间截
		if (second % Female == 0){
			fmt.Println("female")
		}else{
			fmt.Println("man")
		}
		time.Sleep(1000*time.Millisecond)
	}
	
}