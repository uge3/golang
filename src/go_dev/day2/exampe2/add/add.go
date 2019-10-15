package add


import(
	_ "go_dev/day2/exampe2/test"
)

var Name string = "名字"
var Age int = 32
//初始化 init 
func init()  {
	Name = "修改之后"
	Age = 38
}