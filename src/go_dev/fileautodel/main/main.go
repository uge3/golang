package main

import (
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/config" //读取配置文件
)

//读取配置文件
func readconfig() (file_path string, maxsize int) {
	conf, err := config.NewConfig("ini", "./config.conf") //读取一个新的配置文件
	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}
	fmt.Println("conf内容", conf)
	file_path = conf.String("server::file_path")
	fmt.Println("文件路径：", file_path)
	maxsize, err = conf.Int("server::maxsize")
	if err != nil {
		fmt.Println("获取最大文件出错server::maxsize failed err:", err)
		return
	}
	return string(file_path), int(maxsize)
}

func main() {
	var data = ""
	var maxsize = 0
	data, maxsize = readconfig()
	fmt.Println("文件路径+文件名:", data)
	fmt.Printf("最大值：", maxsize)
	deltime := 0 //休眠时间
	for true {
		fmt.Println("停止时间为:", deltime, "秒！")
		time.Sleep(time.Duration(deltime) * time.Second) //进入休眠
		fi, err := os.Stat(data)                         //读取文件信息
		if err == nil {
			fmt.Println("name:", fi.Name())
			fmt.Println("size:", fi.Size())
			fmt.Println("is dir:", fi.IsDir())
			fmt.Println("mode::", fi.Mode())
			fmt.Println("modTime:", fi.ModTime())
		} else if err != nil {
			fmt.Println("查询文件错误:", err)
			fmt.Println(fi)
		}
		if fi != nil {
			fisize := (fi.Size() / 1024) * 1000 //以k为单位
			fmt.Println("文件大小(k):", fisize)
			if int(fisize) > maxsize {
				err := os.Remove(data) //如果超过额定大小，进行删除
				if err != nil {
					fmt.Println("删除失败") // 删除失败
				} else {
					fmt.Println("删除成功") // 删除成功
				}
			} else {
				fmt.Println("文件不够大")
				deltime = int(float64(15000) - float64(fisize)) //对再次删除的时间进行计算，并退出当前循环
				continue
			}
		} else {
			fmt.Println("没有该文件!")
			deltime = 10 //没有文件等待生成
			continue
		}
	}
}
