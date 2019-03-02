package gameModel

import (
	"db"
	"log"
	"constant"
)

func WriteGameEssay(title string,cover string,author string,content string,time string,tag string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.GameEssayInfo{}
	//var args = []interface{}{}
	InsertSql:="insert into blog.game(title,cover,author,content,time,tag) values($1,$2,$3,$4,$5,$6);"
	stmt,err:=db.Db.Prepare(InsertSql)
	if err != nil {
		log.Println("GameModel WriteGame Inserysql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(title,cover,author,content,time,tag)
	if err!=nil{
		log.Println("GameModel WriteGame exce fail")
		return false, err
	}
	//fmt.Println(rows,err)
	//for rows.Next() {
	//	var user constant.LoginInfo
	//	err := rows.Scan(&user.LoginAccount, &user.LoginPassword, &user.Role)
	//	if err != nil {
	//		return nil, err
	//	}
	//	logininfo = append(logininfo,&user)
	//}
	return true,nil
}

func GetAllGameEssay(limit int,offset int)(essayinfo []*constant.GameEssayInfo,essaynumber int,err error){
	var Count int
	CountSql:="select count(*) from blog.game"
	err= db.Db.QueryRow(CountSql).Scan(&Count)
	if err!=nil{
		log.Println("gameModel GetAllGameEssay CountSql exec fail")
		return
	}
	essayinfo=[]*constant.GameEssayInfo{}
	var args = []interface{}{}
	QuerySql:="select id,title,cover,author,content,time,tag from blog.game"
	if limit==-1{
		QuerySql+=";"
	}else{
		QuerySql+=" limit $1 offset $2;"
		args=append(args,limit,offset)
	}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("gameModel GetAllGameEssay QuerySql exec fail")
		//logger.Logger.Error("gameModel GetAllGameEssay QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newessayinfo constant.GameEssayInfo
		err := rows.Scan(&newessayinfo.Id,&newessayinfo.Title,&newessayinfo.Cover,&newessayinfo.Author,&newessayinfo.Content,&newessayinfo.Time,&newessayinfo.Tag)
		if err != nil {
			log.Fatalln(err)
		}
		essayinfo = append(essayinfo,&newessayinfo)
	}
	//fmt.Println(essayinfo)
	//fmt.Println(Count)
	if Count==0{
		return nil,Count,err
	}
	rows.Close()
	return essayinfo,Count,err
}

func GetOneGameEssay(gameid string)(essayinfo []*constant.GameEssayInfo,err error){
	//var Count int
	//CountSql:="select count(*) from essay"
	//err= db.Db.QueryRow(CountSql).Scan(&Count)
	//if err!=nil{
	//	log.Println("gameModel GetAllGameInfo CountSql exec fail")
	//	return
	//}
	essayinfo=[]*constant.GameEssayInfo{}
	var args = []interface{}{}
	QuerySql:="select id,title,cover,author,content,time,tag from blog.game where id = $1"
	args=append(args,gameid)
	//if limit==-1{
	//	QuerySql+=";"
	//}else{
	//	QuerySql+=" limit $1 offset $2;"
	//	args=append(args,limit,offset)
	//}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("gameModel GetOneGameEssay QuerySql exec fail")
		//logger.Logger.Error("gameModel GetOneGameEssay QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newessayinfo constant.GameEssayInfo
		err := rows.Scan(&newessayinfo.Id,&newessayinfo.Title,&newessayinfo.Cover,&newessayinfo.Author,&newessayinfo.Content,&newessayinfo.Time,&newessayinfo.Tag)
		if err != nil {
			log.Fatalln(err)
		}
		essayinfo = append(essayinfo,&newessayinfo)
	}
	if essayinfo==nil{
		return nil,err
	}
	rows.Close()
	return essayinfo,err
}
