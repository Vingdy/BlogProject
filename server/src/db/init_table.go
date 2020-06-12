package db

import (
	"errors"
	"log"
)

//func initAllTable() {
//
//	//登录数据库
//	err := execSQL(`CREATE TABLE IF NOT EXISTS loginuser(
//	id SERIAL NOT NULL,
//	username text NOT NULL,
//	password text NOT NULL,
//	role text NOT NULL,
//	PRIMARY KEY ("id")
//	);`)
//	if err != nil {
//		log.Panicln("init table user failed " + err.Error())
//	}else{
//		log.Println("table loginuser has been created")
//	}
//	//文章数据库
//	err = execSQL(`CREATE TABLE IF NOT EXISTS essay(
//	id SERIAL NOT NULL,
//	title text NOT NULL,
//	author text NOT NULL default '',
//	content text NOT NULL default '',
//	time timestamp NOT NULL default now(),
//	tag text NOT NULL default '',
//	PRIMARY KEY ("id")
//	);`)
//	if err != nil {
//		log.Panicln("init table essay failed " + err.Error())
//	}else{
//		log.Println("table essay has been created")
//	}
//	//游戏数据库
//	err = execSQL(`CREATE TABLE IF NOT EXISTS game(
//	id SERIAL NOT NULL,
//	title text NOT NULL,
//    cover text,
//	author text NOT NULL default '',
//	content text NOT NULL default '',
//	time timestamp NOT NULL default now(),
//	tag text NOT NULL default '',
//	PRIMARY KEY ("id")
//	);`)
//	if err != nil {
//		log.Panicln("init table game failed " + err.Error())
//	}else{
//		log.Println("table game has been created")
//	}
//	err = execSQL(`CREATE TABLE IF NOT EXISTS sentence(
//	id SERIAL NOT NULL,
//	content text NOT NULL default '',
//	time timestamp NOT NULL default now(),
//	PRIMARY KEY ("id")
//	);`)
//	if err != nil {
//		log.Panicln("init table sentence failed " + err.Error())
//	}else{
//		log.Println("table sentence has been created")
//	}
//	err = execSQL(`CREATE TABLE IF NOT EXISTS drawpicture(
//	id SERIAL NOT NULL,
//    title text NOT NULL,
//	src text NOT NULL default '',
//	time timestamp NOT NULL default now(),
//	PRIMARY KEY ("id")
//	);`)
//	if err != nil {
//		log.Panicln("init table drawpicture failed " + err.Error())
//	}else{
//		log.Println("table drawpicture has been created")
//	}
//}
//
//func execSQL(sql string) error {
//	//检测sql语句长度
//	if len(sql) <= 0 {
//		return errors.New("execSQL sql empty")
//	}
//	//sql语句准备
//	stmt, err := Db.Prepare(sql)
//	if err != nil {
//		return err
//	}
//	//事务执行
//	_, err = stmt.Exec()
//	//关闭事务
//	defer stmt.Close()
//	if err != nil {
//		return err
//	}
//	return nil
//}

func initAllTable() {
	//err:=execSQL("CREATE SCHEMA IF NOT EXISTS blog AUTHORIZATION postgre")
	//if err != nil {
	//	log.Panicln("init schema blog failed " + err.Error())
	//}else{
	//	log.Println("schema blog has been created")
	//}
	//登录数据库
	err := execSQL(`CREATE TABLE IF NOT EXISTS blog.loginuser(
	id SERIAL NOT NULL,
	username text NOT NULL,
	password text NOT NULL,
	role text NOT NULL,
	isdelete text NOT NULL default '0',
	PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Panicln("init table blog.loginuser failed " + err.Error())
	} else {
		log.Println("table blog.loginuser has been created")
	}
	//文章数据库
	err = execSQL(`CREATE TABLE IF NOT EXISTS blog.essay(
	id SERIAL NOT NULL,
	title text NOT NULL,
	author text NOT NULL default '',
	content text NOT NULL default '',
	time timestamp NOT NULL default now(),
	tag text NOT NULL default '',
	isdelete text NOT NULL default '0',
	PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Panicln("init table blog.essay failed " + err.Error())
	} else {
		log.Println("table blog.essay has been created")
	}
	//游戏数据库
	err = execSQL(`CREATE TABLE IF NOT EXISTS blog.game(
	id SERIAL NOT NULL,
	title text NOT NULL,
    cover text,
	author text NOT NULL default '',
	content text NOT NULL default '',
	time timestamp NOT NULL default now(),
	tag text NOT NULL default '',
	isdelete text NOT NULL default '0',
	PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Panicln("init table blog.game failed " + err.Error())
	} else {
		log.Println("table blog.game has been created")
	}
	err = execSQL(`CREATE TABLE IF NOT EXISTS blog.sentence(
	id SERIAL NOT NULL,
	content text NOT NULL default '',
	time timestamp NOT NULL default now(),
	tag text NOT NULL default '',
	isdelete text NOT NULL default '0',
	PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Panicln("init table blog.sentence failed " + err.Error())
	} else {
		log.Println("table blog.sentence has been created")
	}
	err = execSQL(`CREATE TABLE IF NOT EXISTS blog.drawpicture(
		id SERIAL NOT NULL,
        title text NOT NULL,
		src text NOT NULL default '',
		time timestamp NOT NULL default now(),
		isdelete text NOT NULL default '0',
		PRIMARY KEY ("id")
		);`)
	if err != nil {
		log.Panicln("init table blog.drawpicture failed " + err.Error())
	} else {
		log.Println("table blog.drawpicture has been created")
	}
	err = execSQL(`CREATE TABLE IF NOT EXISTS blog.userdata(
		id SERIAL NOT NULL,
        name text NOT NULL,
		headpicture text,
		info text NOT NULL,
		PRIMARY KEY ("id")
		);`)
	if err != nil {
		log.Panicln("init table blog.user failed " + err.Error())
	} else {
		log.Println("table blog.user has been created")
	}
}

func execSQL(sql string) error {
	//检测sql语句长度
	if len(sql) <= 0 {
		return errors.New("execSQL sql empty")
	}
	//sql语句准备
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return err
	}
	//事务执行
	_, err = stmt.Exec()
	//关闭事务
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
