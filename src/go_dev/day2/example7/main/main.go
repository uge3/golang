package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Println("the operating system is %s\n", goos)
	var paths = os.Getenv("PATH")
	fmt.Println("path is %s\n", paths)
}
