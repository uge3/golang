package tailf

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/hpcloud/tail"
)

type CollectConf struct {
	LogPath string //日志路径
	Topic   string //主题
} //数组(扩展性)

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
} //根据包的来写，一个对象

type TextMsg struct {
	Msg   string
	Topic string
} //日志信息

type TailObjMgr struct {
	tailObjs []*TailObj

	msgChan chan *TextMsg //日志信息管道

} //管理全部对象

var (
	tailObjMgr *TailObjMgr
) //单例

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
} //接口,读取日志的一行

func InitTall(conf []CollectConf, chanSize int) (err error) {
	fmt.Println("初始化实时跟踪文件变化功能！")

	if len(conf) == 0 {
		err = fmt.Errorf("无效的日志配置文件:%v", conf)
		return
	} //配置文件参数校验
	fmt.Println("配置文件参数校验成功！")

	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	} //单例

	for _, v := range conf {

		obj := &TailObj{
			conf: v,
		}

		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen: true, //重新打开新的文件
			Follow: true, //
			//Location:  &tail.SeekInfo{offset: 0, Whence: 2},//从哪个位置读(保存的位置，重新定位)
			MustExist: false,
			Poll:      true,
		}) //初始化
		if errTail != nil {
			fmt.Println("打开文件失败！tail file err:", errTail)
			err = errTail
			return
		} //打开文件失败！

		obj.tail = tails

		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj) //单例

		go readFromTail(obj) //启动一个协程，真正的去读取日志内容

	} //实例化一个对象

	return
}

func readFromTail(tailObj *TailObj) {

	for true {
		line, ok := <-tailObj.tail.Lines //从管道里读取
		if !ok {
			logs.Warn("管道已经关闭:%s\n", tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue //退出无限循环
		} //管道已经关闭

		textMsg := &TextMsg{
			Msg:   line.Text,          //每行内容
			Topic: tailObj.conf.Topic, //主题
		} //日志内容，结构体

		tailObjMgr.msgChan <- textMsg //读取到日志内容放到管道里

	} //无限循环

}
