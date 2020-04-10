package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Emainl   string `db:"email"`
}

type User struct {
	UserId    int    `json:"user_id"`
	Passwd    string `json:"passwd"`
	Nick      string `json:"nick"`
	Sex       string `json:"sex"`
	Header    string `json:"header"`
	LastLogin string `json:"last_login"`
	Status    int    `json:"status"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	//连接mysql                           用户名        密码   地址                    库
	//database, err := sqlx.Open("mysql", "go_test_user:gotest@tcp(47.107.49.152:3306)/go_test")
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open  mysql failed,", err)
		return
	}
	Db = database
}

//用户结构体

func main() {

	//查询数据                      表名    字段1   ，2 ，3                        数据内容
	var person []Person
	err := Db.Select(&person, "select user_id,sex from persons where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failen", err)
		return
	}
	fmt.Println("select succ:", person)

}
