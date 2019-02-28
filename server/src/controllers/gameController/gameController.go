package gameController

import (
	"net/http"
	"utils/feedback"
	"io/ioutil"
	"log"
	"logger"
	"constant"
	"encoding/json"
	"models/gameModel"
	"net/url"
	"strconv"
)

func WriteGameEssay(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "ReadAll failed:" + err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body获取错误").Response()
		return
	}
	var essayinfo constant.GameEssayInfo
	err = json.Unmarshal(body, &essayinfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		//fb.Response(w, constant.PARA_ERR, "请求body解析json错误", nil)
		return
	}
	gamewrite,err:=gameModel.WriteGameEssay(essayinfo.Title,essayinfo.Author,essayinfo.Content,essayinfo.Time,essayinfo.Tag)
	if err!=nil {
		msg := "gameModel WriteGameEssay run fail:"+err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("WriteGameEssay运行错误").Response()
		return
	}
	if !gamewrite{
		msg:="WriteGame success"
		log.Println(msg)
		logger.Logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("游戏上传失败").Response()
		return
	}
	msg:="WriteGame success"
	log.Println(msg)
	logger.Logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("游戏上传成功").Response()
}

func GetAllGameEssay(w http.ResponseWriter,r *http.Request){
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
	Allgameinfo,AllgameCounter,err:=gameModel.GetAllGameEssay(limitint,offsetint)
	if err!=nil {
		msg := "gameModel GetAllGameEssay run fail:"+err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetAllGameEssay运行错误").Response()
		return
	}
	//fmt.Println(Allgameinfo)
	if AllgameCounter==0{
		msg:="GetAllGameinfo is empty"
		logger.Logger.Info(msg)
		fb.FbCode(constant.FILE_HAS_NOT_EXISTED).FbMsg("游戏列表为空").FbTotal(0).Response()
		return
	}
	msg:="GetAllGameinfo success"
	logger.Logger.Info(msg)
	log.Println(msg)
	//fmt.Println(Allgameinfo)
	fb.FbCode(constant.SUCCESS).FbMsg("游戏列表获取成功").FbData(Allgameinfo).FbTotal(AllgameCounter).Response()
}

func GetOneGameEssay(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	essayid:=queryForm["essayid"][0]
	Onegameinfo,err:=gameModel.GetOneGameEssay(essayid)
	if err!=nil {
		msg := "gameModel GetOneGameEssay run fail:"+err.Error()
		log.Println(msg)
		logger.Logger.Error(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetOneGameEssay运行错误").Response()
		return
	}
	msg:="GetOneGameinfo success"
	logger.Logger.Info(msg)
	log.Println(msg)
	//fmt.Println(Onegameinfo)
	fb.FbCode(constant.SUCCESS).FbMsg("该游戏获取成功").FbData(Onegameinfo).Response()
}


