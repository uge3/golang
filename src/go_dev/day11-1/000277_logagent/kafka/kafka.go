package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer //sarama包的接口，接口不需要指针
)

func Initkafka(addr string) (err error) {

	config := sarama.NewConfig()                              //实例化一个kafka配置
	config.Producer.RequiredAcks = sarama.WaitForAll          //写入kafka,确认写入（SK）
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区（一个队列，随机分配多台机器）
	config.Producer.Return.Successes = true                   //写成功了

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		fmt.Println("初始化kafka失败,必须先保证zookeeper和kafka服务正常启动:", err)
		fmt.Println("准备初始化kafka！！注意配置./conf/logagent.conf（配置文件）//手动创建这个文件")
		fmt.Println(" [kafka] server_addr=192.192.168.60:9092 (本机IP)")

		logs.Error("初始化kafka失败:", err)
		return
	} //kafka生成者的对象,新建一个连接

	//defer client.Close() //函数最后执行，关闭连接

	logs.Debug("初始化kafka成功!")
	return
}

func SendToKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}       //kafka的结构体
	msg.Topic = topic                      //等待匹配的标题
	msg.Value = sarama.StringEncoder(data) //具体信息

	pid, offset, err := client.SendMessage(msg) //发送信息
	if err != nil {
		//fmt.Println("send message failed:", err)
		logs.Error("kafka发送信息失败，错误:%v,信息:%v,主题:%v", err, data, topic)
		return
	} //pid分区(哪台机器)，offset分区里的偏移量(多少M)

	//fmt.Printf("pid:%v offset:%v\n", pid, offset)
	//time.Sleep(10 * time.Millisecond)
	logs.Debug("kafka发送信息成功!pid:%v ,offset:%v ,topic:%v \n", pid, offset, topic)
	return
}

/*
	kafka的安装
	a1)安装JDK,https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html
	a2)配置JDK,https://jingyan.baidu.com/article/6dad5075d1dc40a123e36ea3.html

	b1)安装zookeeper,http://apache.fayea.com/zookeeper/
	b2)配置zookeeper,https://my.oschina.net/demons99/blog/1928362?tdsourcetag=s_pcqq_aiomsg
		3.1、进入D:\zookeeper-3.4.10\conf目录拷贝zoo_sample.cfg文件为zoo.cfg.
        4.1、添加变量 ZOOKEEPER_HOME
                变量值：D:\zookeeper-3.4.10
        4.2、修改变量 path
                添加内容：;%ZOOKEEPER_HOME%/bin;%ZOOKEEPER_HOME%/conf
	b3)启动zookeeper,D:\gorelated\zookeeper\bin>zkServer.cmd

	c1)安装kafka,http://kafka.apache.org/downloads
	c2)配置kafka,https://blog.csdn.net/evankaka/article/details/52421314
	c3)打开config目录下的server.properties,修改log.dirs为D:\kafka_logs,修改advertised.host.name=192.192.168.60(服务器IP)
	c4)启动kafka,D:\gorelated\kafka>.\bin\windows\kafka-server-start.bat .\config\server.properties

	c5)创建主题,D:\gorelated\kafka>.\bin\windows\kafka-topics.bat --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic nginx_log
	c6)创建生产者，D:\gorelated\kafka>.\bin\windows\kafka-console-producer.bat --broker-list localhost:9092 --topic nginx_log
	c7)创建消费者（新版本），D:\gorelated\kafka>.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic nginx_log --from-beginning

*/

/*
	go kafka的安装与测试
	自动安装:
	go git的安装与测试
	https://git-scm.com/downloads
	https://www.cnblogs.com/osfipin/p/4891610.html
	D:\\project> go get github.com/Shopify/sarama

	手动安装：
	https://github.com/Shopify/sarama
	下载后，在 D:\project\src 下创建  /github.com/Shopify/sarama
	将解压后的文件全部复制到 D:/project/src/github.com/Shopify/sarama

	修复插件问题1：config.go->"golang.org/x/net/proxy"
	https://blog.csdn.net/xie1xiao1jun/article/details/79421136
	在 D:\project\src 下创建  \golang.org\x\
	终端(命令行)切换(cd)到 D:\project\src\golang.org\x
	git clone https://github.com/golang/net.git net
	go install net

	修复插件问题2：手动下载复制到 D:/project/src/github.com/pierrec/lz4
	https://github.com/pierrec/lz4
	github.com/DataDog/zstd

	修复插件问题3：MinGW64 离线安装包V8.1
	http://www.3h3.com/soft/124832.html
	http://211.143.146.242/1Q2W3E4R5T6Y7U8I9O0P1Z2X3C4V5B/soft1.3h3.com/3H3mingw.64Bit.rar
*/
