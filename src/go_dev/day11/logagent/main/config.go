package main

import (
	"errors"
	"fmt"
	"go_dev/day11/logagent/tailf"

	"github.com/astaxie/beego/config"
)

var (
	appConfig *Config //全局变量
)

type Config struct {
	logLevel  string
	logPath   string
	chanSize  int
	kafkaAddr string

	collectConf []tailf.CollectConf //用数组
}

func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf //读 取配置内容 到变量
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}//读取日志路径失败

	cc.Topic = conf.String("collect::topic")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}//读取主题失败
	
	appConfig.collectConf = append(appConfig.collectConf, cc) // 放入数组
	return
}

//加载
func loadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed err :", err)
		return
	}
	appConfig = &Config{}                               //配置实例
	appConfig.logLevel = conf.String("logs::log_level") //日志级别
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug" //默认的日志级别
	}

	appConfig.logPath = conf.String("logs::log_path") //日志路径
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}
	appConfig.chanSize, err = conf.Int("collect::chan_size") //管道配置
	if err != nil {
		appConfig.chanSize = 100
	}
	appConfig.kafkaAddr = conf.String("kafka::server_aeer") //读取kafka配置
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed,err :%v\n", err)
		return
	}
	return

}
