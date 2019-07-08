package sentenceController

import (
	"net/http"
	"utils/feedback"
	"io/ioutil"
	"constant"
	"encoding/json"
	"models/sentenceModel"
	"net/url"
	"strconv"
	"logger"
)

func WriteSentence(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "ReadAll failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body获取错误").Response()
		return
	}
	var essayinfo constant.SentenceInfo
	err = json.Unmarshal(body, &essayinfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		//fb.Response(w, constant.PARA_ERR, "请求body解析json错误", nil)
		return
	}
	sentencewrite,err:=sentenceModel.WriteSentence(essayinfo.Content,essayinfo.Time)
	if err!=nil {
		msg := "sentenceModel WriteSentence run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("WriteSentence运行错误").Response()
		return
	}
	if !sentencewrite{
		msg:="WriteSentence success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("句子上传失败").Response()
		return
	}
	msg:="WriteSentence success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("句子上传成功").Response()
}

func GetAllSentence(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	//member:=queryForm["member"][0]
	offset:=queryForm["offset"][0]
	limit:=queryForm["limit"][0]
	searchstring:=queryForm["searchstring"][0]
	limitint, err := strconv.Atoi(limit)
	if err!=nil{
		msg:="limit to int failed:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("limit转换错误").Response()
		//data:=model.FeedBackErrorHandle(501,msg)
		//fmt.Fprintln(w,string(data))
		return
	}
	offsetint,err:=strconv.Atoi(offset)
	if err!=nil{
		msg:="offset to int failed:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		//data:=model.FeedBackErrorH
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("offset转换错误").Response()
		//fmt.Fprintln(w,string(data))
		return
	}
	Allsentenceinfo,AllsentenceCounter,err:=sentenceModel.GetAllSentenceInfo(limitint,offsetint,searchstring)
	if err!=nil {
		msg := "sentenceModel GetAllSentenceInfo run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetAllSentenceInfo运行错误").Response()
		return
	}
	if AllsentenceCounter==0{
		msg:="GetAllSentenceinfo is empty"
		//logger.Logger.Info(msg)
		//log.Println(msg)
		logger.Info(msg)
		fb.FbCode(constant.FILE_HAS_NOT_EXISTED).FbMsg("句子列表为空").FbTotal(0).Response()
		return
	}
	msg:="GetAllSentenceinfo success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("句子列表获取成功").FbData(Allsentenceinfo).FbTotal(AllsentenceCounter).Response()
}

func GetOneSentence(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	sentenceid:=queryForm["sentenceid"][0]
	Onesentenceinfo,err:=sentenceModel.GetOneSentence(sentenceid)
	if err!=nil {
		msg := "sentenceModel GetOneSentence run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetOneSentence运行错误").Response()
		return
	}
	msg:="GetOneSentence success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该句子获取成功").FbData(Onesentenceinfo).Response()
}

func GetSentenceTime(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	//queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	//essayid:=queryForm["sreach"][0]
	sentencetag,err:=sentenceModel.GetSentenceTime()
	if err!=nil {
		msg := "sentenceModel GetSentenceTime run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetSentenceTime运行错误").Response()
		return
	}
	msg:="GetSentenceTime success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("句子时间归档获取成功").FbData(sentencetag).Response()
}

func UpdateSentence(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "ReadAll failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body获取错误").Response()
		return
	}
	var sentenceinfo constant.SentenceInfo
	err = json.Unmarshal(body, &sentenceinfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		//fb.Response(w, constant.PARA_ERR, "请求body解析json错误", nil)
		return
	}
	id, err := strconv.Atoi(sentenceinfo.Id)
	if err!=nil{
		msg:="essayinfo.Id to int failed:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("essayinfo.Id转换错误").Response()
		//data:=model.FeedBackErrorHandle(501,msg)
		//fmt.Fprintln(w,string(data))
		return
	}
	sentenceok,err:=sentenceModel.UpdateSentence(sentenceinfo.Content,sentenceinfo.Time,id)
	if err!=nil {
		msg := "sentenceModel UpdateSentence run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("UpdateSentence运行错误").Response()
		return
	}
	if !sentenceok{
		msg:="UpdateSentence success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("该句子修改失败").Response()
		return
	}
	msg:="UpdateSentence success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该句子修改成功").Response()
}

func DeleteSentence(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "ReadAll failed:" + err.Error()
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body获取错误").Response()
		return
	}
	var essayinfo map[string]interface{}
	err = json.Unmarshal(body, &essayinfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		return
	}
	essayid,ok:=essayinfo["essayid"].(string)
	if !ok{
		msg := "get map key failed:" + err.Error()
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("发送body中不存在key值").Response()
		return
	}
	sentenceinfo,err:=sentenceModel.DeleteSentence(essayid)
	if err!=nil {
		msg := "sentenceModel DeleteSentence run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("DeleteSentence运行错误").Response()
		return
	}
	if !sentenceinfo{
		msg:="DeleteSentence success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("该句子删除失败").Response()
		return
	}
	msg:="DeleteSentence success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该句子删除成功").Response()
}
