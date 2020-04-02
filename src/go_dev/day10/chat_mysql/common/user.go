package common

const (
	UserStatusOnline  = 1    //表示在线
	UserStatusOffline = iota //表示离线
)

//用户结构体
type User struct {
	UserId    int    `json:"user_id"`
	Passwd    string `json:"passwd"`
	Nick      string `json:"nick"`
	Sex       string `json:"sex"`
	Header    string `json:"header"`
	LastLogin string `json:"last_login"`
	Status    int    `json:"status"`
}
