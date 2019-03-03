package db

//连接数据库初始化

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

var Db *sql.DB//持续连接

func InitDB(host, port, user, pwd, dbName, driverName string) {
	log.Println("=====启动连接数据库=====")
	//构建连接字符串
	dateSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbName)
	fmt.Println(dateSource)
	db, err := sql.Open(driverName, dateSource)
	if err != nil {
		log.Panicln(err)
}
	Db = db
	err = Db.Ping()
	if err != nil {
		log.Println("InitDB failed at Ping " + err.Error())
		log.Panicln(err)
	}
	db.SetMaxOpenConns(5)
	//初始化table
	initAllTable()
}
