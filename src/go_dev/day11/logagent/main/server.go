package main

import (
	"fmt"
	"go_dev/day11/logagent/kafka"
	"go_dev/day11/logagent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

func serverRun() (err error) {
	fmt.Println("========5555555555555555====")
	for {
		msg := tailf.GetOneLine()
		fmt.Println("========444444444444444444444====")
		err = sendTokafka(msg) //写入kafka
		if err != nil {
			logs.Error("send to kafka failed err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}

	return
}

//写入kafka
func sendTokafka(msg *tailf.TextMsg) (err error) {
	// logs.Debug("read msg:%s,topic:%s", msg.Msg, msg.Topic)
	fmt.Printf("read msg:%s,topic:%s", msg.Msg, msg.Topic)
	err = kafka.SendTokafka(msg.Msg, msg.Topic)
	return
}
