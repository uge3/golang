package main

import "go_dev/day10/chat_mysql/chat_server/model"

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	// mgr = model.NewUserMgr(pool)
	mgr = model.NewUserMgr(Db)
}
