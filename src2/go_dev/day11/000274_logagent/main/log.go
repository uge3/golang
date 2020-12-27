package main //固定写法，定义成go包

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug //7
	case "warn":
		return logs.LevelWarn 
	case "info":
		return logs.LevelInfo //6
	case "trace":
		return logs.LevelTrace 
	}
	return logs.LevelDebug //7

} //字符串转数字

func initLogger() (err error) {
	config := make(map[string]interface{}) //数组

	config["filename"] = appConfig.logPath                //文件路径
	config["level"] = convertLogLevel(appConfig.logLevel) //日志级别（等级）

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("初始化错误,序列化错误:", err)
		return
	} //序列化成JSON字符串，数组

	logs.SetLogger(logs.AdapterFile, string(configStr)) //JSON字符串，数组
	return
}
