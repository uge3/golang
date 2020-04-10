package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	//连接mysql                           用户名        密码   地址                    库
	// database, err := sqlx.Open("mysql", "go_test_user:gotest@tcp(47.107.49.152:3306)/go_test")
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open  mysql failed;", err)
		return
	}
	Db = database
}

func main() {

	//initRedis("127.0.0.1:6379", 16, 1004, time.Second*300) //初始化redis
	//initRedis("47.107.49.152:6379","admin123456" ,16, 1004, time.Second*300) //初始化redis
	initUserMgr()              //初始化用户
	runServer("0.0.0.0:10000") //运行服务端

}
