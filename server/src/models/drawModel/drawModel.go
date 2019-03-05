package drawModel

import (
	"db"
	"log"
	"constant"
)

func WriteDrawPicture(title string,src string,time string,tag string)(drawpictureinsertok bool, err error){
	//drawpictureinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	InsertSql:="insert into blog.drawpicture(title,src,time,tag) values($1,$2,$3,$4);"
	stmt,err:=db.Db.Prepare(InsertSql)
	if err != nil {
		log.Println("DrawModel WriteDrawPicture Inserysql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(title,src,time,tag)
	if err!=nil{
		log.Println("DrawModel WriteDrawPicture exce fail")
		return false, err
	}
	return true,nil
}

func GetAllDrawPicture(limit int,offset int,searchstring string)(drawpictureinfo []*constant.DrawPictureInfo,drawpicturenumber int,err error){
	var Count int
	var args = []interface{}{}
	CountSql:="select count(*) from blog.drawpicture"
	CountSql+=" where (to_char(time,'yyyy-mm-dd hh24:mi:ss') like $1 or tag like $2 or title like $3)"
	args=append(args,"%"+searchstring+"%","%"+searchstring+"%","%"+searchstring+"%")
	CountSql+=" and isdelete = '0'"
	err= db.Db.QueryRow(CountSql,args...).Scan(&Count)
	if err!=nil{
		log.Println("drawModel GetAllDrawPicture CountSql exec fail")
		return
	}
	drawpictureinfo=[]*constant.DrawPictureInfo{}
	args = []interface{}{}
	QuerySql:="select id,title,src,time,tag from blog.drawpicture"
	QuerySql+=" where (to_char(time,'yyyy-mm-dd hh24:mi:ss') like $1 or tag like $2 or title like $3)"
	args=append(args,"%"+searchstring+"%","%"+searchstring+"%","%"+searchstring+"%")
	QuerySql+=" and isdelete = '0'"
	QuerySql+=" order by time desc"
	if limit==-1{
		QuerySql+=";"
	}else{
		QuerySql+=" limit $4 offset $5;"
		args=append(args,limit,offset)
	}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("drawModel GetAllDrawPicture QuerySql exec fail")
		//logger.Logger.Error("drawModel GetAllDrawPicture QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newdrawpictureinfo constant.DrawPictureInfo
		err := rows.Scan(&newdrawpictureinfo.Id,&newdrawpictureinfo.Title,&newdrawpictureinfo.Src,&newdrawpictureinfo.Time,&newdrawpictureinfo.Tag)
		if err != nil {
			log.Fatalln(err)
		}
		drawpictureinfo = append(drawpictureinfo,&newdrawpictureinfo)
	}
	if Count==0{
		return nil,Count,err
	}
	rows.Close()
	return drawpictureinfo,Count,err
}

func GetOneDrawPicture(drawid string)(drawpictureinfo []*constant.DrawPictureInfo,err error){
	//var Count int
	//CountSql:="select count(*) from drawpicture"
	//err= db.Db.QueryRow(CountSql).Scan(&Count)
	//if err!=nil{
	//	log.Println("drawModel GetAllDrawInfo CountSql exec fail")
	//	return
	//}
	drawpictureinfo=[]*constant.DrawPictureInfo{}
	var args = []interface{}{}
	QuerySql:="select id,title,src,time,tag from blog.drawpicture where id = $1"
	args=append(args,drawid)
	//if limit==-1{
	//	QuerySql+=";"
	//}else{
	//	QuerySql+=" limit $1 offset $2;"
	//	args=append(args,limit,offset)
	//}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("drawModel GetOneDrawPicture QuerySql exec fail")
		//logger.Logger.Error("drawModel GetOneDrawPicture QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newdrawpictureinfo constant.DrawPictureInfo
		err := rows.Scan(&newdrawpictureinfo.Id,&newdrawpictureinfo.Title,&newdrawpictureinfo.Src,&newdrawpictureinfo.Time,&newdrawpictureinfo.Tag)
		if err != nil {
			log.Fatalln(err)
		}
		drawpictureinfo = append(drawpictureinfo,&newdrawpictureinfo)
	}
	if drawpictureinfo==nil{
		return nil,err
	}
	rows.Close()
	return drawpictureinfo,err
}

func GetDrawPictureTag()(taginfo []*constant.TagInfo,err error){
	taginfo=[]*constant.TagInfo{}
	var args = []interface{}{}
	QuerySql:="SELECT tag,COUNT(tag) FROM blog.drawpicture where isdelete = '0' GROUP BY tag"
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("drawModel GetDrawPictureTag QuerySql exec fail")
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

func GetDrawPictureTime()(taginfo []*constant.TagInfo,err error){
	taginfo=[]*constant.TagInfo{}
	var args = []interface{}{}
	QuerySql:="SELECT to_char(time,'yyyy-mm'),COUNT(to_char(time,'yyyy-mm')) FROM blog.drawpicture where isdelete = '0' GROUP BY to_char(time,'yyyy-mm')"
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("drawModel GetDrawPictureTime QuerySql exec fail")
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

func UpdateDrawPicture(title string,src string,time string,tag string,id int)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	UpdateSql:="update blog.drawpicture set title=$1,src=$2,time=$3,tag=$4 where id=$5;"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("drawModel UpdateDrawPicture Updatesql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(title,src,time,tag,id)
	if err!=nil{
		log.Println("drawModel UpdateDrawPicture exce fail")
		return false, err
	}
	return true,nil
}

func DeleteDrawPicture(id string)(essayinsertok bool, err error){
	//essayinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	UpdateSql:="update blog.drawpicture set isdelete='1' where id=$1;"
	stmt,err:=db.Db.Prepare(UpdateSql)
	if err != nil {
		log.Println("drawModel DeleteDrawPicture DeleteSql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(id)
	if err!=nil{
		log.Println("drawModel DeleteDrawPicture exce fail")
		return false, err
	}
	return true,nil
}