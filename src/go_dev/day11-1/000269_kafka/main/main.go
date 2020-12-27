package main

import (
	"fmt"
	"net"
	"time"

	"github.com/Shopify/sarama"
)

func GetAddres() string {
	ip := "127.0.0.1"

	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.获取接口失败, err:", err.Error())
		return ip
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						//fmt.Println(ipnet.IP.String())
						ip = ipnet.IP.String()
						return ip
					}
				}
			}
		}
	}
	return ip
}

func main() { //固定写法，程序的入口
	//kafka
	fmt.Println("必须先保证zookeeper和kafka服务正常启动!")
	fmt.Println(".\\bin\\windows\\kafka-server-start.bat .\\config\\server.properties")

	config := sarama.NewConfig()                              //实例化一个kafka配置
	config.Producer.RequiredAcks = sarama.WaitForAll          //写入kafka,确认写入（SK）
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区（一个队列，随机分配多台机器）
	config.Producer.Return.Successes = true                   //写成功了

	//client, err := sarama.NewSyncProducer([]string{"192.192.168.60:9092"}, config) //服务IP
	ip := GetAddres()
	fmt.Println("本机IP:", ip)
	ip_port := fmt.Sprintf("%s:9092", ip)
	fmt.Println("本机IP和端口:", ip_port)
	client, err := sarama.NewSyncProducer([]string{ip_port}, config) //服务IP
	if err != nil {
		fmt.Println("生产者已经关闭，错误:", err)
		fmt.Println("注意配置!kafka\\config\\server.properties文件，添加(本机IP地址)advertised.host.name=127.0.0.1")
		return
	} //kafka生成者的对象,新建一个连接
	defer client.Close() //函数最后执行，关闭连接

	n := 1
	for {
		n++
		msg := &sarama.ProducerMessage{}                                     //kafka的结构体
		msg.Topic = "nginx_log"                                              //等待匹配的标题
		msg.Value = sarama.StringEncoder(fmt.Sprintf("test,这是好的测试内容,%v", n)) //具体信息

		pid, offset, err := client.SendMessage(msg) //发送信息
		if err != nil {
			fmt.Println("发送信息失败:", err)
			return
		} //pid分区(哪台机器)，offset分区里的偏移量(多少M)

		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(10 * time.Millisecond)
	}
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

/*
	b3)启动zookeeper,D:\gorelated\zookeeper\bin>zkServer.cmd
	c4)启动kafka,D:\gorelated\kafka>.\bin\windows\kafka-server-start.bat .\config\server.properties
	必须先保证zookeeper和kafka服务正常启动

	fmt.Println("------运行方法,在终端cd(切换)到当前目录：------")
	fmt.Println("cd D:/project/src/go_dev/day11/000269_kafka/main")
	fmt.Println("go run main.go ")

	fmt.Println("编译成二进制的运行方法：D:\\project>  ")
	fmt.Println("go build go_dev/day11/000269_kafka/main")
	fmt.Println("注意配置环境变量  新建用户变量 变量名(N)GOPATH  变量值(V)D:\\project")
	fmt.Println("在终端输入：D:\\project>  main.exe")
	fmt.Println("------退出程序！------")

	查看消费者
	c7)创建消费者（新版本），D:\gorelated\kafka>.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic nginx_log --from-beginning

*/
