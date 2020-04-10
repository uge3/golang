package model

import (
	"encoding/json"
	"fmt"
	"go_dev/day10/chat_mysql/common"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	UserTable = "users"
)

//用户管理
type UserMgr struct {
	db *sqlx.DB
}

//用户 mysql 连接
func NewUserMgr(db *sqlx.DB) (mgr *UserMgr) {
	mgr = &UserMgr{
		db: db,
	}
	return
}

//用户验证
func (p *UserMgr) getUser(id int) (user *common.User, err error) {
	fmt.Println("用户验证中...................")
	var users []common.Users
	user = &common.User{}
	fmt.Println(users)
	err = p.db.Select(&users, "select user_id, passwd from persons where user_id=?", id)
	fmt.Println("数据提取完成 =", id, users[0].UserId)
	if err != nil {
		//	fmt.Println("===", users)
		fmt.Println("exec failen用户验证出错:", err)
		return
	}
	fmt.Println("user内容：", user)
	fmt.Println("select succ:", user)
	user.UserId = users[0].UserId
	user.Passwd = users[0].Passwd
	return
}

//登录
func (p *UserMgr) Login(id int, passwd string) (user *common.User, err error) {
	user, err = p.getUser(id)
	if err != nil {
		return
	}

	if user.UserId != id || user.Passwd != passwd {
		err = ErrInvalidPasswd
		return
	}
	user.Status = common.UserStatusOnline
	user.LastLogin = fmt.Sprintf("%v", time.Now()) //上线时间
	return
}

//用户注册
func (p *UserMgr) Register(user *common.User) (err error) {
	// conn := p.db.Select()
	// defer conn.Close()

	fmt.Println("[]", user.UserId)
	if user == nil {
		fmt.Println("invalild,user")
		err = ErrInvalidParams
		return
	}
	fmt.Println("用户注册中。。。")
	_, err = p.getUser(int(user.UserId)) //如果用户存在
	if err == nil {
		err = ErrUserExist //返回
		return
	}
	fmt.Println("用户检测后。。。")
	if err != ErrUserNotExist { //如果错非用户 存在，说明有其它问题
		return
	}
	fmt.Println("+++++**************++++++")
	data, err := json.Marshal(user) //注册成功 ，存起来
	fmt.Println("+++++++++++++++++", string(data))
	fmt.Println("_______________", user)
	fmt.Println(user.UserId, user.Passwd)

	if err != nil {
		return
	}
	_, err = p.db.Exec("insert into persons(user_id,passwd,nick,sex,header,last_login,status)values(?,?,?,?,?,?,?)", user.UserId, user.Passwd, user.Nick, user.Sex, user.Header, user.LastLogin, user.Status)
	if err != nil {
		return
	}
	return
}
