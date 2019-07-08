package drawController

import (
	"constant"
	"net/http"
	"utils/feedback"
	"net/url"
	"models/drawModel"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"logger"
)

func WriteDrawPicture(w http.ResponseWriter,r *http.Request){
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
	var drawinfo constant.DrawPictureInfo
	err = json.Unmarshal(body, &drawinfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		//fb.Response(w, constant.PARA_ERR, "请求body解析json错误", nil)
		return
	}
	drawwrite,err:=drawModel.WriteDrawPicture(drawinfo.Title,drawinfo.Src,drawinfo.Time,drawinfo.Tag)
	if err!=nil {
		msg := "drawModel WriteDrawPicture run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("WriteDrawPicture运行错误").Response()
		return
	}
	if !drawwrite{
		msg:="WriteDraw success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("摸鱼图上传失败").Response()
		return
	}
	msg:="WriteDraw success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("摸鱼图上传成功").Response()
}

func GetAllDrawPicture(w http.ResponseWriter,r *http.Request){
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
		//logger.Logger.Error(msg)
		//log.Println(msg)
		//data:=model.FeedBackErrorH
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("offset转换错误").Response()
		//fmt.Fprintln(w,string(data))
		return
	}
	Alldrawinfo,AlldrawCounter,err:=drawModel.GetAllDrawPicture(limitint,offsetint,searchstring)
	if err!=nil {
		msg := "drawModel GetAllDrawPicture run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetAllDrawPicture运行错误").Response()
		return
	}
	//fmt.Println(Alldrawinfo)
	if AlldrawCounter==0{
		msg:="GetAllDrawPictureinfo is empty"
		//logger.Logger.Info(msg)
		//log.Println(msg)
		logger.Info(msg)
		fb.FbCode(constant.FILE_HAS_NOT_EXISTED).FbMsg("摸鱼图列表为空").FbTotal(0).Response()
		return
	}
	msg:="GetAllDrawinfo success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("摸鱼图列表获取成功").FbData(Alldrawinfo).FbTotal(AlldrawCounter).Response()
}

func GetOneDrawPicture(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	essayid:=queryForm["pictureid"][0]
	Onedrawinfo,err:=drawModel.GetOneDrawPicture(essayid)
	if err!=nil {
		msg := "drawModel GetOneDrawPicture run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetOneDrawPicture运行错误").Response()
		return
	}
	msg:="GetOneDrawinfo success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该摸鱼图获取成功").FbData(Onedrawinfo).Response()
}

func GetDrawPictureTag(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	//queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	//essayid:=queryForm["sreach"][0]
	drawtag,err:=drawModel.GetDrawPictureTag()
	if err!=nil {
		msg := "drawModel GetDrawPictureTag run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetDrawPictureTag运行错误").Response()
		return
	}
	msg:="GetDrawPictureTag success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("摸鱼标签归档获取成功").FbData(drawtag).Response()
}

func GetDrawPictureTime(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	//queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	//essayid:=queryForm["sreach"][0]
	drawtag,err:=drawModel.GetDrawPictureTime()
	if err!=nil {
		msg := "drawModel GetDrawPictureTime run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetDrawPictureTime运行错误").Response()
		return
	}
	msg:="GetDrawPictureTime success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("摸鱼时间归档获取成功").FbData(drawtag).Response()
}

func UpdateDrawPicture(w http.ResponseWriter,r *http.Request){
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
	var drawinfo constant.DrawPictureInfo
	err = json.Unmarshal(body, &drawinfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		//fb.Response(w, constant.PARA_ERR, "请求body解析json错误", nil)
		return
	}
	id, err := strconv.Atoi(drawinfo.Id)
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
	drawok,err:=drawModel.UpdateDrawPicture(drawinfo.Title,drawinfo.Src,drawinfo.Time,drawinfo.Tag,id)
	if err!=nil {
		msg := "drawModel UpdateDrawPicturerun fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("UpdateDrawPicture运行错误").Response()
		return
	}
	if !drawok{
		msg:="UpdateDrawPicture success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("该摸鱼图修改失败").Response()
		return
	}
	msg:="UpdateDrawPicture success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该摸鱼图修改成功").Response()
}

func DeleteDrawPicture(w http.ResponseWriter,r *http.Request){
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
	drawinfo,err:=drawModel.DeleteDrawPicture(essayid)
	if err!=nil {
		msg := "drawModel DeletePicture run fail:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("DeleteDrawPicture运行错误").Response()
		return
	}
	if !drawinfo{
		msg:="DeleteDrawPicture success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("该摸鱼图删除失败").Response()
		return
	}
	msg:="DeleteDrawPicture success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该摸鱼图删除成功").Response()
}