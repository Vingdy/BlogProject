package blogModel

import (
	"db"
	"log"
	"constant"
	"logger"
)

func WriteBlogEssay(title string,author string,content string,time string,tag string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	InsertSql:="insert into essay(title,author,content,time,tag) values($1,$2,$3,$4,$5);"
	stmt,err:=db.Db.Prepare(InsertSql)
	if err != nil {
		log.Println("BlogModel WriteBlogEssay Inserysql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(title,author,content,time,tag)
	if err!=nil{
		log.Println("BlogModel WriteBlogEssay exce fail")
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

func GetAllBlogEssay(limit int,offset int)(essayinfo []*constant.BlogEssayInfo,essaynumber int,err error){
	var Count int
	CountSql:="select count(*) from essay"
	err= db.Db.QueryRow(CountSql).Scan(&Count)
	if err!=nil{
		log.Println("blogModel GetAllBlogEssay CountSql exec fail")
		return
	}
	essayinfo=[]*constant.BlogEssayInfo{}
	var args = []interface{}{}
	QuerySql:="select id,title,author,content,time,tag from essay"
	if limit==-1{
		QuerySql+=";"
	}else{
		QuerySql+=" limit $1 offset $2;"
		args=append(args,limit,offset)
	}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("blogModel GetAllBlogEssay QuerySql exec fail")
		logger.Logger.Error("blogModel GetAllBlogEssay QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newessayinfo constant.BlogEssayInfo
		err := rows.Scan(&newessayinfo.Id,&newessayinfo.Title,&newessayinfo.Author,&newessayinfo.Content,&newessayinfo.Time,&newessayinfo.Tag)
		if err != nil {
			log.Fatalln(err)
		}
		essayinfo = append(essayinfo,&newessayinfo)
	}
	if Count==0{
		return nil,Count,err
	}
	rows.Close()
	return essayinfo,Count,err
}

func GetOneBlogEssay(blogid string)(essayinfo []*constant.BlogEssayInfo,err error){
	//var Count int
	//CountSql:="select count(*) from essay"
	//err= db.Db.QueryRow(CountSql).Scan(&Count)
	//if err!=nil{
	//	log.Println("blogModel GetAllBlogInfo CountSql exec fail")
	//	return
	//}
	essayinfo=[]*constant.BlogEssayInfo{}
	var args = []interface{}{}
	QuerySql:="select id,title,author,content,time,tag from essay where id = $1"
	args=append(args,blogid)
	//if limit==-1{
	//	QuerySql+=";"
	//}else{
	//	QuerySql+=" limit $1 offset $2;"
	//	args=append(args,limit,offset)
	//}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("blogModel GetOneBlogEssay QuerySql exec fail")
		logger.Logger.Error("blogModel GetOneBlogEssay QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newessayinfo constant.BlogEssayInfo
		err := rows.Scan(&newessayinfo.Id,&newessayinfo.Title,&newessayinfo.Author,&newessayinfo.Content,&newessayinfo.Time,&newessayinfo.Tag)
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

func GetBlogEssayTag(){

}