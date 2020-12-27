package main

import (
	"go_dev/day11-1/000277_logagent/kafka"
	"go_dev/day11-1/000277_logagent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

func serverRun() (err error) {

	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("发送到kafka失败，错误：%v", err)
			time.Sleep(time.Second)
			continue
		}
	}

	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {

	//fmt.Printf("读取信息:%s,读取主题:%s \n", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)

	return
}
