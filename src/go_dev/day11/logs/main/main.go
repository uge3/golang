package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs" //日志管理
)

func main() {
	config := make(map[string]interface{})       //初始化配置文件句柄
	config["filename"] = "./logs/logcollect.log" //配置文件名 路径
	config["level"] = logs.LevelDebug            //日志提示等级
	configStr, err := json.Marshal(config)       //json格式化
	if err != nil {
		fmt.Println("marshal failed err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr)) //日志输出到内 容

	logs.Debug("this is a test ,my name is  %s", "stu01")
	logs.Trace("this is a  trace ,my name is %s", "str01")
	logs.Warn("this is a warn ,my name is  %s", "stu03")
	fmt.Println("日志输出完成!:")
}
