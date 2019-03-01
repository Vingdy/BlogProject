package drawController

import (
	"logger"
	"constant"
	"log"
	"net/http"
	"utils/feedback"
	"net/url"
	"models/drawModel"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

func WriteDrawPicture(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "ReadAll failed:" + err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body获取错误").Response()
		return
	}
	var drawinfo constant.DrawPictureInfo
	err = json.Unmarshal(body, &drawinfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		//fb.Response(w, constant.PARA_ERR, "请求body解析json错误", nil)
		return
	}
	drawwrite,err:=drawModel.WriteDrawPicture(drawinfo.Title,drawinfo.Src,drawinfo.Time)
	if err!=nil {
		msg := "drawModel WriteDrawPicture run fail:"+err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("WriteDrawPicture运行错误").Response()
		return
	}
	if !drawwrite{
		msg:="WriteDraw success"
		log.Println(msg)
		logger.Logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("摸鱼图上传失败").Response()
		return
	}
	msg:="WriteDraw success"
	log.Println(msg)
	logger.Logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("摸鱼图上传成功").Response()
}

func GetAllDrawPicture(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	//member:=queryForm["member"][0]
	offset:=queryForm["offset"][0]
	limit:=queryForm["limit"][0]
	limitint, err := strconv.Atoi(limit)
	if err!=nil{
		msg:="limit to int failed:"+err.Error()
		logger.Logger.Error(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("limit转换错误").Response()
		//data:=model.FeedBackErrorHandle(501,msg)
		//fmt.Fprintln(w,string(data))
		return
	}
	offsetint,err:=strconv.Atoi(offset)
	if err!=nil{
		msg:="offset to int failed:"+err.Error()
		logger.Logger.Error(msg)
		//data:=model.FeedBackErrorH
		fb.FbCode(constant.PARA_ERR).FbMsg("offset转换错误").Response()
		//fmt.Fprintln(w,string(data))
		return
	}
	Alldrawinfo,AlldrawCounter,err:=drawModel.GetAllDrawPicture(limitint,offsetint)
	if err!=nil {
		msg := "drawModel GetAllDrawPicture run fail:"+err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetAllDrawPicture运行错误").Response()
		return
	}
	//fmt.Println(Alldrawinfo)
	if AlldrawCounter==0{
		msg:="GetAllBloginfo is empty"
		logger.Logger.Info(msg)
		log.Println(msg)
		fb.FbCode(constant.FILE_HAS_NOT_EXISTED).FbMsg("摸鱼图列表为空").FbTotal(0).Response()
		return
	}
	msg:="GetAllDrawinfo success"
	log.Println(msg)
	logger.Logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("摸鱼图列表获取成功").FbData(Alldrawinfo).FbTotal(AlldrawCounter).Response()
}

func GetOneDrawPicture(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	essayid:=queryForm["pictureid"][0]
	Onedrawinfo,err:=drawModel.GetOneDrawPicture(essayid)
	if err!=nil {
		msg := "drawModel GetOneDrawPicture run fail:"+err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetOneDrawPicture运行错误").Response()
		return
	}
	msg:="GetOneDrawinfo success"
	logger.Logger.Info(msg)
	log.Println(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该摸鱼图获取成功").FbData(Onedrawinfo).Response()
}