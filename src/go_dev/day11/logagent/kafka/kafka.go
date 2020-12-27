package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer //接口
)

//连接KAFKA
func InitKafka(addr string) (err error) {
	//初始化配置
	config := sarama.NewConfig()                              //生产者模式
	config.Producer.RequiredAcks = sarama.WaitForAll          //
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区负载均衡
	config.Producer.Return.Successes = true                   //返回

	client, err = sarama.NewSyncProducer([]string{addr}, config) //生产都的对象
	if err != nil {
		fmt.Println("producer close,err:", err)
		logs.Error("init kafka producer fialed err,", err)
		return
	}
	logs.Debug("init kafka succ")
	return
}

func SendTokafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test,my message is  good")

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed:", err)
		logs.Error("send message  fialed err:%v ,data%v ,topic:%v", err, data, topic) //发送失败
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	logs.Debug("send succ,pid:%v offset:%v,topic:%v\n", pid, offset, topic)
	return

}
