package main //固定写法，定义成go包

import ( //引入各种包

	"fmt"
	"net"
)

func LocalIPv4s() []string {

	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ips
	}

	for _, a := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())

		}
	}
	return ips

}

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

	LocalIPs := LocalIPv4s()
	fmt.Println(LocalIPs)

	addr := GetAddres()
	fmt.Println(addr)

}

/*
	fmt.Println("------运行方法,在终端cd(切换)到当前目录：------")
	fmt.Println("cd D:\\project\\src\\go_dev\\day2\\000045_exampe1\\main")
	fmt.Println("go run main.go ")

	fmt.Println("编译成二进制的运行方法：D:\\project>  ")
	fmt.Println("go build go_dev\\day2\\000045_exampe1\\main")
	fmt.Println("注意配置环境变量  新建用户变量 变量名(N)GOPATH  变量值(V)D:\\project")
	fmt.Println("在终端输入：D:\\project>  main.exe")
	fmt.Println("------退出程序！------")
*/
