package loginModel

import (
	"log"
	"database/sql"
	"db"
	"constant"
)

var QuerySql string

func CheckLoginAcc(username string)(isexist bool,err error) {
	var haveusername string
	QuerySql = "select username from blog.loginuser where username = $1;"
	//stmt,err:=init_db.Db.Prepare(QuerySql)
	//if err!=nil{
	//	log.Println("LoginModel CheckLoginAcc sql prepare fail")
	//	return false,err
	//}
	//defer stmt.Close()
	err = db.Db.QueryRow(QuerySql, username).Scan(&haveusername)
	//fmt.Println(haveAcc,"+",err)
	//if err!=nil{
	//	log.Println("loginModel CheckLoginAcc sql query fail")
	//	return false, err
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("loginModel CheckLoginAcc not account found")
			//logger.Logger.Error("loginModel CheckLoginAcc not account found")
			return false,nil
		} else {
			log.Println("loginModel CheckLoginAcc Querysql query fail"+err.Error())
			//logger.Logger.Error("loginModel CheckLoginAcc Querysql query fail"+err.Error())
			return false,err
		}
	}
	return true,nil
}

func Login(username string,password string)(logininfo []*constant.LoginInfo, err error) {
	logininfo=[]*constant.LoginInfo{}
	//var args = []interface{}{}
	QuerySql = "select username,password,role from blog.loginuser where username = $1 and password = $2;"
	stmt, err := db.Db.Prepare(QuerySql)
	if err != nil {
		log.Println("loginModel Login Querysql prepare fail")
		//logger.Logger.Error("loginModel Login Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username, password)
	if err!=nil{
		log.Println("loginModel Login Querysql query fail")
		//logger.Logger.Error("loginModel Login Querysql query fail")
		return nil, err
	}
	//fmt.Println(rows,err)
	for rows.Next() {
		var user constant.LoginInfo
		err := rows.Scan(&user.LoginAccount, &user.LoginPassword, &user.Role)
		if err != nil {
			return nil, err
		}
		logininfo = append(logininfo,&user)
	}
	return logininfo,nil
}