package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	Id   int
	Name string
	Age  int
}

var db *sqlx.DB

func initDB() (err error) {
	dsn := "go_user:231231@tcp(127.0.0.1:3306)/go_base?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	fmt.Println("数据库访问成功")
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func queryRowDemo() {
	sqlstr := "select * from user where id = ?"
	var u user
	err := db.Get(&u, sqlstr, 1)
	if err != nil {
		fmt.Printf("query failed with err :%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

// 查询多条数据示例，真的很方便hhh
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db error with err : %v\n", err)
	}
	queryRowDemo()
	queryMultiRowDemo()
}
