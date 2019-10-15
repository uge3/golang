package main

import "fmt"

func main() {
	i := 0
HERE:
	print(i)
	fmt.Printf("\n=======\n")
	i++
	if i == 5 {
		return
	}
	goto HERE
}
