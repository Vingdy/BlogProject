package drawModel

import (
	"db"
	"log"
	"logger"
	"constant"
)

func WriteDrawPicture(title string,src string,time string)(drawpictureinsertok bool, err error){
	//drawpictureinfo=[]*constant.EssayInfo{}
	//var args = []interface{}{}
	InsertSql:="insert into drawpicture(title,src,time) values($1,$2,$3);"
	stmt,err:=db.Db.Prepare(InsertSql)
	if err != nil {
		log.Println("DrawModel WriteDrawPicture Inserysql prepare fail")
		return false, err
	}
	defer stmt.Close()
	_,err = stmt.Exec(title,src,time)
	if err!=nil{
		log.Println("DrawModel WriteDrawPicture exce fail")
		return false, err
	}
	return true,nil
}

func GetAllDrawPicture(limit int,offset int)(drawpictureinfo []*constant.DrawPictureInfo,drawpicturenumber int,err error){
	var Count int
	CountSql:="select count(*) from drawpicture"
	err= db.Db.QueryRow(CountSql).Scan(&Count)
	if err!=nil{
		log.Println("drawModel GetAllDrawPicture CountSql exec fail")
		return
	}
	drawpictureinfo=[]*constant.DrawPictureInfo{}
	var args = []interface{}{}
	QuerySql:="select id,title,src,time from drawpicture"
	if limit==-1{
		QuerySql+=";"
	}else{
		QuerySql+=" limit $1 offset $2;"
		args=append(args,limit,offset)
	}
	rows,err:=db.Db.Query(QuerySql,args...)
	if err!=nil{
		log.Println("drawModel GetAllDrawPicture QuerySql exec fail")
		logger.Logger.Error("drawModel GetAllDrawPicture QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newdrawpictureinfo constant.DrawPictureInfo
		err := rows.Scan(&newdrawpictureinfo.Id,&newdrawpictureinfo.Title,&newdrawpictureinfo.Src,&newdrawpictureinfo.Time)
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
	QuerySql:="select id,title,src,time from drawpicture where id = $1"
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
		logger.Logger.Error("drawModel GetOneDrawPicture QuerySql exec fail")
		return
	}
	for rows.Next() {
		//fmt.Println("have")
		var newdrawpictureinfo constant.DrawPictureInfo
		err := rows.Scan(&newdrawpictureinfo.Id,&newdrawpictureinfo.Title,&newdrawpictureinfo.Src,&newdrawpictureinfo.Time)
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

func GetDrawPictureTag(){

}