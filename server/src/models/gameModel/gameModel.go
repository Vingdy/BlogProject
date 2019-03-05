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

func GetAllGameEssay(limit int,offset int,searchstring string)(essayinfo []*constant.GameEssayInfo,essaynumber int,err error){
	var Count int
	var args = []interface{}{}
	CountSql:="select count(*) from blog.game"
	CountSql+=" where (to_char(time,'yyyy-mm-dd hh24:mi:ss') like $1 or tag like $2 or title like $3 or author like $4)"
	args=append(args,"%"+searchstring+"%","%"+searchstring+"%","%"+searchstring+"%","%"+searchstring+"%")
	CountSql+=" and isdelete = '0'"
	err= db.Db.QueryRow(CountSql,args...).Scan(&Count)
	if err!=nil{
		log.Println("gameModel GetAllGameEssay CountSql exec fail")
		return
	}
	essayinfo=[]*constant.GameEssayInfo{}
	args = []interface{}{}
	QuerySql:="select id,title,cover,author,content,time,tag from blog.game"
	QuerySql+=" where (to_char(time,'yyyy-mm-dd hh24:mi:ss') like $1 or tag like $2 or title like $3 or author like $4)"
	args=append(args,"%"+searchstring+"%","%"+searchstring+"%","%"+searchstring+"%","%"+searchstring+"%")
	QuerySql+=" and isdelete = '0'"
	QuerySql+=" order by time desc"
	if limit==-1{
		QuerySql+=";"
	}else{
		QuerySql+=" limit $5 offset $6;"
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

func GetGameEssayTag()(taginfo []*constant.TagInfo,err error){
	taginfo=[]*constant.TagInfo{}
	var args = []interface{}{}
	QuerySql:="SELECT tag,COUNT(tag) FROM blog.game where isdelete = '0' GROUP BY tag"
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("gameModel GetGameEssayTag QuerySql exec fail")
		return
	}
	for rows.Next() {
		var newtaginfo constant.TagInfo
		err := rows.Scan(&newtaginfo.Name,&newtaginfo.Number)
		if err != nil {
			log.Fatalln(err)
		}
		taginfo = append(taginfo,&newtaginfo)
	}
	rows.Close()
	return taginfo,err
}

func GetGameEssayTime()(taginfo []*constant.TagInfo,err error){
	taginfo=[]*constant.TagInfo{}
	var args = []interface{}{}
	QuerySql:="SELECT to_char(time,'yyyy-mm'),COUNT(to_char(time,'yyyy-mm')) FROM blog.game where isdelete = '0' GROUP BY to_char(time,'yyyy-mm')"
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("gameModel GetGameEssayTime QuerySql exec fail")
		return
	}
	for rows.Next() {
		var newtaginfo constant.TagInfo
		err := rows.Scan(&newtaginfo.Name,&newtaginfo.Number)
		if err != nil {
			log.Fatalln(err)
		}
		taginfo = append(taginfo,&newtaginfo)
	}
	rows.Close()
	return taginfo,err
}

func UpdateGameEssay(title string,cover string,author string,content string,time string,tag string,id int)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	UpdateSql:="update blog.game set title=$1,cover=$2,author=$3,content=$4,time=$5,tag=$6 where id=$7;"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("gameModel UpdateGameEssay Updatesql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(title,cover,author,content,time,tag,id)
	if err!=nil{
		log.Println("gameModel UpdateGameEssay exce fail")
		return false, err
	}
	return true,nil
}

func DeleteGameEssay(id string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	UpdateSql:="update blog.game set isdelete='1' where id=$1;"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("gameModel DeleteGameEssay DeleteSql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(id)
	if err!=nil{
		log.Println("gameModel DeleteGameEssay exce fail")
		return false, err
	}
	return true,nil
}