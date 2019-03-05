package userModel

import (
	"constant"
	"db"
	"log"
)

func GetUserData()(userinfo []*constant.UserInfo,err error){
	userinfo=[]*constant.UserInfo{}
	var args = []interface{}{}
	QuerySql:="SELECT name,headpicture,info from blog.userdata"
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("userModel GetUserData QuerySql exec fail")
		return
	}
	for rows.Next() {
		var newuserinfo constant.UserInfo
		err := rows.Scan(&newuserinfo.Name,&newuserinfo.Headpicture,&newuserinfo.Info)
		if err != nil {
			log.Fatalln(err)
		}
		userinfo = append(userinfo,&newuserinfo)
	}
	rows.Close()
	return userinfo,err
}

func UpdateUserData(name string,headpicture string,info string)(updateuserok bool,err error){
	//userinfo=[]*constant.UserInfo{}
	var args = []interface{}{}
	UpdateSql:="update blog.userdata set name=$1,headpicture=$2,info=$3 where id=1"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("userModel UpdateUserData Updatesql prepare fail")
		return false, err
	}
	args=append(args,name,headpicture,info)
	defer stmt.Close()
	_,err = stmt.Exec(args...)
	if err!=nil{
		log.Println("userModel UpdateUserData exce fail")
		return false, err
	}
	return true,nil
}

