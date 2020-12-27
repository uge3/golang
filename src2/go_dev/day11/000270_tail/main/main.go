package main //固定写法，定义成go包
import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() { //固定写法，程序的入口
	//tail

	//注意换行方式
	filename := "./txt/s1.txt" //s1.txt,s2.txt,s3.txt

	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true, //重新打开新的文件
		Follow: true, //
		//Location:  &tail.SeekInfo{offset: 0, Whence: 2},//从哪个位置读(保存的位置，重新定位)
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("打开文件失败！tail file err:", err)
		return
	} //打开文件失败！

	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines //从管道里读取
		if !ok {
			fmt.Printf("管道已经关闭tail file close reopen,filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue //退出无限循环
		} //管道已经关闭

		fmt.Println("信息:", msg)
	}

}

/*
	go tail的安装与测试
	自动安装:
	go git的安装与测试
	https://git-scm.com/downloads
	https://www.cnblogs.com/osfipin/p/4891610.html

	D:\\project> go get github.com/hpcloud/tail

	手动安装：
	https://github.com/hpcloud/tail
	下载后，在 D:\project\src 下创建  /github.com/hpcloud/tail
	将解压后的文件全部复制到 D:/project/src/github.com/hpcloud/tail
*/

/*
	fmt.Println("------运行方法,在终端cd(切换)到当前目录：------")
	fmt.Println("注意配置环境变量  新建用户变量 变量名(N)GOPATH  变量值(V)D:/project")

	fmt.Println("编译成二进制的运行方法：d:/project/src/go_dev/day11/000270_tail> ")
	fmt.Println("go build go_dev/day11/000270_tail/main")
	fmt.Println("在终端输入：d:/project/src/go_dev/day11/000270_tail> main.exe ")
	fmt.Println("------退出程序！------")
*/
