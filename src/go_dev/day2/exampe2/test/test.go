package test

import("fmt")

var Name string ="test"
var Age int = 1

func init(){
	fmt.Println("this is test package")
	fmt.Println("test.package.Name=",Name)

	Age=12
	fmt.Println("test.package chang Age = ",Age)

}