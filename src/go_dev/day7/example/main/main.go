package main

import (
	"fmt"
	"go_dev/day7/example/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	//模拟实现主机实例
	var insts ([]*balance.Instance)
	for i := 0; i < 3; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	// var balancer balance.Balancer
	// var conf = "random" //默认
	// if len(os.Args) > 1 {
	// 	conf = os.Args[1] //接收命令行参数,1表示接收第一个参数
	// }
	// if conf == "random" {
	// 	balancer = &balance.RandomBalance{} //随机轮巡
	// } else if conf == "ronudrobin" {
	// 	balancer = &balance.RoundRobinBalance{} //随机轮巡
	// }

	//balancer := &balance.RandomBalance{} //随机轮巡

	var balanceName = "random" //默认
	if len(os.Args) > 1 {
		balanceName = os.Args[1] //接收命令行参数,1表示接收第一个参数
	}

	for {
		//inst, err := balancer.DoBalance(insts)
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
