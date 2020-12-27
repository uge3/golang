package main

import (
	"fmt"

	sarama "github.com/Shopify/sarama"
)

func main() {
	//初始化配置
	config := sarama.NewConfig()                              //生产者模式
	config.Producer.RequiredAcks = sarama.WaitForAll          //
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区负载均衡
	config.Producer.Return.Successes = true                   //返回

	clinet, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config) //生产都的对象
	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}
	defer clinet.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test,my message is  good")

	pid, offset, err := clinet.SendMessage(msg) //发送到KAFKA
	if err != nil {
		fmt.Println("send message failed:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

}
