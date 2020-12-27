package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//日志级别转 换
func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
func initLogger() (err error) {
	config := make(map[string]interface{}) //初始化配置文件句柄
	config["filename"] = appConfig.logPath //配置文件名 路径
	config["level"] = logs.LevelDebug      //日志提示等级
	configStr, err := json.Marshal(config) //json格式化
	if err != nil {
		fmt.Println("initLogger failed ,marshal err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr)) //日志输出到内 容

	return
}
