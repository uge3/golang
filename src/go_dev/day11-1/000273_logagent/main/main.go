package main //固定写法，定义成go包
import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() { //固定写法，程序的入口

	filename := "./conf/logagent.conf" //配置文件路径
	err := loadConf("ini", filename)   //加载初始化配置文件
	if err != nil {
		fmt.Printf("加载配置文件错误: %v \n", err)
		panic("panic加载配置文件错误")
		return
	} //加载初始化配置文件
	fmt.Println("加载初始化配置文件成功！！")

	err = initLogger() //初始化日志
	if err != nil {
		fmt.Printf("初始化日志失败:%v \n", err)
		panic("panic初始化日志失败")
		return
	} //初始化日志
	fmt.Println("初始化日志成功！！")

	logs.Debug("Debug初始化日志")
	logs.Debug("Debug加载配置文件:%v \n", appConfig)

	fmt.Println("查看日志文件 log_path=./logs/logagent.log")
}

/*
	fmt.Println("------运行方法,在终端cd(切换)到当前目录：------")
	fmt.Println("注意配置环境变量  新建用户变量 变量名(N)GOPATH  变量值(V)D:/project")

	fmt.Println("编译成二进制的运行方法：d:/project/src/go_dev/day11/000273_logagent> ")
	fmt.Println("go build go_dev/day11/000273_logagent/main")
	fmt.Println("在终端输入：d:/project/src/go_dev/day11/000273_logagent> main.exe ")
	fmt.Println("------退出程序！------")

*/
