package sentenceModel

import (
	"log"
	"db"
	"constant"
	"logger"
)

func WriteSentence(content string,time string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.SentenceInfo{}
	//var args = []interface{}{}
	InsertSql:="insert into sentence(content,time) values($1,$2);"
	stmt,err:=db.Db.Prepare(InsertSql)
	if err != nil {
		log.Println("SentenceModel WriteSentence Inserysql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(content,time)
	if err!=nil{
		log.Println("SentenceModel WriteSentence exce fail")
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

func GetAllSentenceInfo(limit int,offset int)(essayinfo []*constant.SentenceInfo,essaynumber int,err error){
	var Count int
	CountSql:="select count(*) from sentence"
	err= db.Db.QueryRow(CountSql).Scan(&Count)
	if err!=nil{
		log.Println("sentenceModel GetAllSentenceInfo CountSql exec fail")
		return
	}
	essayinfo=[]*constant.SentenceInfo{}
	var args = []interface{}{}
	QuerySql:="select id,content,time from sentence"
	if limit==-1{
		QuerySql+=";"
	}else{
		QuerySql+=" limit $1 offset $2;"
		args=append(args,limit,offset)
	}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("sentenceModel GetAllSentenceInfo QuerySql exec fail")
		logger.Logger.Error("sentenceModel GetAllSentenceInfo QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newsentenceinfo constant.SentenceInfo
		err := rows.Scan(&newsentenceinfo.Id,&newsentenceinfo.Content,&newsentenceinfo.Time,)
		if err != nil {
			log.Fatalln(err)
		}
		essayinfo = append(essayinfo,&newsentenceinfo)
	}
	if essayinfo==nil{
		return nil,Count,err
	}
	rows.Close()
	return essayinfo,Count,err
}