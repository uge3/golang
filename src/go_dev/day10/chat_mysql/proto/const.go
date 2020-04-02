package proto

//在线统计，注册
const (
	UserLogin           = "user_login"
	UserLoginRes        = "user_login_res"
	UserRegister        = "user_register"
	UserStatusNotifyRes = "user_statusnotify" //上下线状态
	UserSendMessageCmd  = "user_send_message"  //用户发消息
	UserRecvMessageCmd  = "user_recv_message"//用户收消息
)

const (
	UserOnline  = 1 //上线
	UserOffline = 2 // 下线
)
