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

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	//连接mysql                           用户名        密码   地址                    库
	database, err := sqlx.Open("mysql", "go_test_user:gotest@tcp(47.107.49.152:3306)/go_test")
	if err != nil {
		fmt.Println("open  mysql failed,", err)
		return
	}
	Db = database
}
func main() {
	//更新数据                      表名    字段1   ，2 ，3                        数据内容

	n, err := Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failen", err)
		return
	}
	fmt.Println("delete:", &n)
}
