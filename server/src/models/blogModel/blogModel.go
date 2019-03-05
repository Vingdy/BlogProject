package blogModel

import (
	"db"
	"log"
	"constant"
)

func WriteBlogEssay(title string,author string,content string,time string,tag string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	InsertSql:="insert into blog.essay(title,author,content,time,tag) values($1,$2,$3,$4,$5);"
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

func GetAllBlogEssay(limit int,offset int,searchstring string)(essayinfo []*constant.BlogEssayInfo,essaynumber int,err error){
	var Count int
	var args = []interface{}{}
	CountSql:="select count(*) from blog.essay"
	CountSql+=" where (to_char(time,'yyyy-mm-dd hh24:mi:ss') like $1 or tag like $2 or title like $3 or author like $4)"
	args=append(args,"%"+searchstring+"%","%"+searchstring+"%","%"+searchstring+"%","%"+searchstring+"%")
	CountSql+=" and isdelete = '0'"
	err= db.Db.QueryRow(CountSql,args...).Scan(&Count)
	if err!=nil{
		log.Println("blogModel GetAllBlogEssay CountSql exec fail")
		return
	}
	essayinfo=[]*constant.BlogEssayInfo{}
	args = []interface{}{}
	QuerySql:="select id,title,author,content,time,tag from blog.essay"
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
		log.Println("blogModel GetAllBlogEssay QuerySql exec fail")
		//logger.Logger.Error("blogModel GetAllBlogEssay QuerySql exec fail")
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
	QuerySql:="select id,title,author,content,time,tag from blog.essay where id = $1"
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
		//logger.Logger.Error("blogModel GetOneBlogEssay QuerySql exec fail")
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

func GetBlogEssayTag()(taginfo []*constant.TagInfo,err error){
	taginfo=[]*constant.TagInfo{}
	var args = []interface{}{}
	QuerySql:="SELECT tag,COUNT(tag) FROM blog.essay where isdelete = '0' GROUP BY tag"
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("blogModel GetBlogEssayTag QuerySql exec fail")
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

func GetBlogEssayTime()(taginfo []*constant.TagInfo,err error){
	taginfo=[]*constant.TagInfo{}
	var args = []interface{}{}
	QuerySql:="SELECT to_char(time,'yyyy-mm'),COUNT(to_char(time,'yyyy-mm')) FROM blog.essay where isdelete = '0' GROUP BY to_char(time,'yyyy-mm')"
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("blogModel GetBlogEssayTime QuerySql exec fail")
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

func UpdateBlogEssay(title string,author string,content string,time string,tag string,id int)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	UpdateSql:="update blog.essay set title=$1,author=$2,content=$3,time=$4,tag=$5 where id=$6;"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("blogModel UpdateBlogEssay Updatesql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(title,author,content,time,tag,id)
	if err!=nil{
		log.Println("blogModel UpdateBlogEssay exce fail")
		return false, err
	}
	return true,nil
}

func DeleteBlogEssay(id string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	UpdateSql:="update blog.essay set isdelete='1' where id=$1;"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("blogModel DeleteBlogEssay DeleteSql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(id)
	if err!=nil{
		log.Println("blogModel DeleteBlogEssay exce fail")
		return false, err
	}
	return true,nil
}