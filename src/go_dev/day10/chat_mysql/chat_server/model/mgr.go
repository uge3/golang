package model

import (
	"encoding/json"
	"fmt"
	"go_dev/day10/chat_mysql/common"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	UserTable = "users"
)

//用户管理
type UserMgr struct {
	pool *redis.Pool
}

//用户 redis 连接池
func NewUserMgr(pool *redis.Pool) (mgr *UserMgr) {
	mgr = &UserMgr{
		pool: pool,
	}
	return
}

//用户验证
func (p *UserMgr) getUser(conn redis.Conn, id int) (user *common.User, err error) {
	result, err := redis.String(conn.Do("HGet", UserTable, fmt.Sprintf("%d", id)))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrUserNotExist
		}
		return
	}
	user = &common.User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return
	}
	return
}

//登录
func (p *UserMgr) Login(id int, passwd string) (user *common.User, err error) {
	conn := p.pool.Get()
	defer conn.Close()
	user, err = p.getUser(conn, id)
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
	conn := p.pool.Get()
	defer conn.Close()
	if user == nil {
		fmt.Println("invalild,user")
		err = ErrInvalidParams
		return
	}
	_, err = p.getUser(conn, user.UserId) //如果用户存在
	if err == nil {
		err = ErrUserExist //返回
		return
	}
	if err != ErrUserNotExist { //如果错非用户 存在，说明有其它问题
		return
	}
	data, err := json.Marshal(user) //注册成功 ，存起来
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", UserTable, fmt.Sprintf("%d", user.UserId), string(data))
	if err != nil {
		return
	}
	return
}
