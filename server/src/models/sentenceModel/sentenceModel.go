package sentenceModel

import (
	"log"
	"db"
	"constant"
)

func WriteSentence(content string,time string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.SentenceInfo{}
	//var args = []interface{}{}
	InsertSql:="insert into blog.sentence(content,time) values($1,$2);"
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

func GetAllSentenceInfo(limit int,offset int,searchstring string)(essayinfo []*constant.SentenceInfo,essaynumber int,err error){
	var Count int
	var args = []interface{}{}
	CountSql:="select count(*) from blog.sentence"
	CountSql+=" where (to_char(time,'yyyy-mm-dd hh24:mi:ss') like $1 or content like $2)"
	args=append(args,"%"+searchstring+"%","%"+searchstring+"%")
	CountSql+=" and isdelete = '0'"
	err= db.Db.QueryRow(CountSql,args...).Scan(&Count)
	if err!=nil{
		log.Println("sentenceModel GetAllSentenceInfo CountSql exec fail")
		return
	}
	essayinfo=[]*constant.SentenceInfo{}
	args = []interface{}{}
	QuerySql:="select id,content,time from blog.sentence"
	QuerySql+=" where (to_char(time,'yyyy-mm-dd hh24:mi:ss') like $1 or content like $2)"
	args=append(args,"%"+searchstring+"%","%"+searchstring+"%")
	QuerySql+=" and isdelete = '0'"
	QuerySql+=" order by time desc"
	if limit==-1{
		QuerySql+=";"
	}else{
		QuerySql+=" limit $3 offset $4;"
		args=append(args,limit,offset)
	}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("sentenceModel GetAllSentenceInfo QuerySql exec fail")
		//logger.Logger.Error("sentenceModel GetAllSentenceInfo QuerySql exec fail")
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
	if Count==0{
		return nil,Count,err
	}
	rows.Close()
	return essayinfo,Count,err
}

func GetOneSentence(sentenceid string)(sentenceinfo []*constant.SentenceInfo,err error){
	sentenceinfo=[]*constant.SentenceInfo{}
	var args = []interface{}{}
	QuerySql:="select id,content,time from blog.sentence where id = $1"
	args=append(args,sentenceid)
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("sentenceModel GetOneSentence QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newsentenceinfo constant.SentenceInfo
		err := rows.Scan(&newsentenceinfo.Id,&newsentenceinfo.Content,&newsentenceinfo.Time)
		if err != nil {
			log.Fatalln(err)
		}
		sentenceinfo = append(sentenceinfo,&newsentenceinfo)
	}
	if sentenceinfo==nil{
		return nil,err
	}
	rows.Close()
	return sentenceinfo,err
}

func GetSentenceTime()(taginfo []*constant.TagInfo,err error){
	taginfo=[]*constant.TagInfo{}
	var args = []interface{}{}
	QuerySql:="SELECT to_char(time,'yyyy-mm'),COUNT(to_char(time,'yyyy-mm')) FROM blog.sentence where isdelete = '0' GROUP BY to_char(time,'yyyy-mm')"
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("sentenceModel GetSentenceTime QuerySql exec fail")
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

func UpdateSentence(content string,time string,id int)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	UpdateSql:="update blog.sentence set content=$1,time=$2 where id=$3;"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("sentenceModel UpdateSentence Updatesql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(content,time,id)
	if err!=nil{
		log.Println("sentenceModel UpdateSentence exce fail")
		return false, err
	}
	return true,nil
}

func DeleteSentence(id string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	UpdateSql:="update blog.sentence set isdelete='1' where id=$1;"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("sentenceModel DeleteSentence DeleteSql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(id)
	if err!=nil{
		log.Println("sentenceModel DeleteSentence exce fail")
		return false, err
	}
	return true,nil
}