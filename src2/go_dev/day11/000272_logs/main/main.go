package main //固定写法，定义成go包
import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() { //固定写法，程序的入口
	//logs 日志库

	config := make(map[string]interface{}) //数组

	config["filename"] = "./logs/logagent.log" //文件路径
	config["level"] = logs.LevelDebug          //日志级别（等级）
	//config["level"] = logs.LevelInfo //日志级别（等级）

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("序列化错误:", err)
		return
	} //序列化成JSON字符串，数组

	logs.SetLogger(logs.AdapterFile, string(configStr)) //JSON字符串，数组

	logs.Debug("这是测试模式 %s", "stu01")
	logs.Trace("这是跟踪模式 %s", "stu02")
	logs.Warn("这是警告模式 %s", "stu03")
	fmt.Println("查看日志文件 config[“filename”] = “./logs/logagent.log” //文件路径")
}

/*
	go logs的安装与测试

	手动安装：
	https://github.com/astaxie/beego/
	下载后，在 D:\project\src 下创建  /github.com/astaxie/beego/
	将解压后的文件全部复制到 D:/project/src/github.com/astaxie/beego/

	补丁：  github.com/shiena/ansicolor
	自动安装:
	go git的安装与测试
	https://git-scm.com/downloads
	https://www.cnblogs.com/osfipin/p/4891610.html

	D:\\project> go get github.com/shiena/ansicolor
*/

/*
	fmt.Println("------运行方法,在终端cd(切换)到当前目录：------")
	fmt.Println("注意配置环境变量  新建用户变量 变量名(N)GOPATH  变量值(V)D:/project")

	fmt.Println("编译成二进制的运行方法：d:/project/src/go_dev/day11/000272_logs> ")
	fmt.Println("go build go_dev/day11/000272_logs/main")
	fmt.Println("在终端输入：d:/project/src/go_dev/day11/000272_logs> main.exe ")
	fmt.Println("------退出程序！------")
*/
