package tailf

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

type CollectConf struct { //收集配置
	LogPath string //路径
	Topic   string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

//消息结构体  用于管道
type TextMsg struct {
	Msg   string
	Topic string
}

//消息管理
type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg // 管道
}

var (
	tailObjMgr *TailObjMgr //创建一个单例
)

//获得管道内数据
func GetOneLine() (msg *TextMsg) {
	fmt.Println("========666666====")
	msg = <-tailObjMgr.msgChan
	fmt.Println("========7777777====")
	return
}

func InitTail(conf []CollectConf, chanSize int) (err error) {
	fmt.Println("初始化实时跟踪文件变化功能！")
	if len(conf) == 0 {
		err = fmt.Errorf("invalid config for log collect, conf:%v", conf)
		return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize), //初始化管道
	} //配置管理
	//遍历配置文件
	fmt.Println("初始化完成。。。。。。")
	for _, v := range conf {
		obj := &TailObj{
			conf: v,
		}

		tails, errTail := tail.TailFile(v.LogPath, tail.Config{ //初始化
			ReOpen: true, //重新打开新的文件
			Follow: true,
			//Location:  &tail.SeekInfo{offset: 0, Whence: 2},//从哪个位置读(保存的位置，重新定位)
			MustExist: false,
			Poll:      true,
		})
		if errTail != nil {
			err = errTail
			fmt.Println("tail file err :", err)
			return
		}
		obj.tail = tails
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj) //放入单例
		fmt.Println("tailobjmgr:", tails)

		go readFromTail(obj)
	}

	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		msgLine, ok := <-tailObj.tail.Lines //消息行 从管道取出
		fmt.Println("消息行 从管道取出。。。")
		if !ok {
			fmt.Printf("tail file close reopen,filename:%s\n", tailObj.tail.Filename)
			logs.Warn("tail file close reopen,filename:%s\n", tailObj.tail.Filename) //写入日志
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msgLine)
		textMsg := &TextMsg{
			Msg:   msgLine.Text, //消息内容
			Topic: tailObj.conf.Topic,
		}
		tailObjMgr.msgChan <- textMsg //放入管道
	}

}
