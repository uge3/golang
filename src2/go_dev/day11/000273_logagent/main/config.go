package main

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/config"
)

var (
	appConfig *Config
) //全局变量

type Config struct {
	logLevel string
	logPath  string

	collectConf []CollectConf //数组(扩展性)

} //存读取配置文件的信息

type CollectConf struct {
	LogPath string
	Topic   string
} //数组(扩展性)

func loadCollectConf(conf config.Configer) (err error) { //多态接口

	var cc CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	} //读取日志路径失败

	cc.Topic = conf.String("collect::topic")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::topic")
		return
	} //读取主题失败

	appConfig.collectConf = append(appConfig.collectConf, cc) //全局变量
	return
} //读取日志相关的信息

func loadConf(confType, filename string) (err error) {

	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("加载初始化配置文件失败:", err)
		return
	} //找不到路径文件

	appConfig = &Config{} //生成实例
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	} //日志级别,如果读取不到配置，就使用默认配置

	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	} //日志路径,如果读取不到配置，就使用默认配置

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("收集 加载初始化配置文件失败: %v \n", err)
		return
	} //读取日志相关的信息

	return

}

/* ./conf/logagent.conf（配置文件）//手动创建这个文件
[logs]
log_level=debug
log_path=./logs/logagent.log

[collect]
log_path=/home/work/logs/nginx/access.log
topic=nginx_log
*/
