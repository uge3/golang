package tailf

import "fmt"

type CollectConf struct {
	LogPath string
	Topic   string
} //数组(扩展性)

func InitTall(conf []CollectConf) (err error) {
	fmt.Println("初始化实时跟踪文件变化功能！")
	return
}
