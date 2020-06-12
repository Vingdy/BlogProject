package blogController

import (
	"constant"
	"encoding/json"
	"io/ioutil"
	"logger"
	"models/blogModel"
	"net/http"
	"net/url"
	"strconv"
	"utils/feedback"
)

func WriteBlogEssay(w http.ResponseWriter, r *http.Request) {
	fb := feedback.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "ReadAll failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body获取错误").Response()
		return
	}
	var essayinfo constant.BlogEssayInfo
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
	blogwrite, err := blogModel.WriteBlogEssay(essayinfo.Title, essayinfo.Author, essayinfo.Content, essayinfo.Time, essayinfo.Tag)
	if err != nil {
		msg := "blogModel WriteBlogEssay run fail:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("WriteBlogEssay运行错误").Response()
		return
	}
	if !blogwrite {
		msg := "WriteBlog success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("博客上传失败").Response()
		return
	}
	msg := "WriteBlog success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("博客上传成功").Response()
}

func GetAllBlogEssay(w http.ResponseWriter, r *http.Request) {
	//logger.Info("test")
	fb := feedback.NewFeedBack(w)
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	//member:=queryForm["member"][0]
	offset := queryForm["offset"][0]
	limit := queryForm["limit"][0]
	searchstring := queryForm["searchstring"][0]
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		msg := "limit to int failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("limit转换错误").Response()
		//data:=model.FeedBackErrorHandle(501,msg)
		//fmt.Fprintln(w,string(data))
		return
	}
	offsetint, err := strconv.Atoi(offset)
	if err != nil {
		msg := "offset to int failed:" + err.Error()
		//logger.Logger.Error(msg)
		//log.Println(msg)
		//data:=model.FeedBackErrorH
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("offset转换错误").Response()
		//fmt.Fprintln(w,string(data))
		return
	}
	Allbloginfo, AllblogCounter, err := blogModel.GetAllBlogEssay(limitint, offsetint, searchstring)
	if err != nil {
		msg := "blogModel GetAllBlogEssay run fail:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetAllBlogEssay运行错误").Response()
		return
	}
	//fmt.Println(Allbloginfo)
	if AllblogCounter == 0 {
		msg := "GetAllBloginfo is empty"
		//logger.Logger.Info(msg)
		//log.Println(msg)
		logger.Info(msg)
		fb.FbCode(constant.FILE_HAS_NOT_EXISTED).FbMsg("博客列表为空").FbTotal(0).Response()
		return
	}
	msg := "GetAllBloginfo success"
	//log.Println(msg)

	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("博客列表获取成功").FbData(Allbloginfo).FbTotal(AllblogCounter).Response()
}

func GetOneBlogEssay(w http.ResponseWriter, r *http.Request) {
	fb := feedback.NewFeedBack(w)
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	essayid := queryForm["essayid"][0]
	Onebloginfo, err := blogModel.GetOneBlogEssay(essayid)
	if err != nil {
		msg := "blogModel GetOneBlogEssay run fail:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetOneBlogEssay运行错误").Response()
		return
	}
	msg := "GetOneBloginfo success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该博客获取成功").FbData(Onebloginfo).Response()
}

func GetBlogEssayTag(w http.ResponseWriter, r *http.Request) {
	fb := feedback.NewFeedBack(w)
	//queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	//essayid:=queryForm["sreach"][0]
	blogtag, err := blogModel.GetBlogEssayTag()
	if err != nil {
		msg := "blogModel GetBlogEssayTag run fail:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetBlogEssayTag运行错误").Response()
		return
	}
	msg := "GetBlogTag success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("博客标签归档获取成功").FbData(blogtag).Response()
}

func GetBlogEssayTime(w http.ResponseWriter, r *http.Request) {
	fb := feedback.NewFeedBack(w)
	//queryForm,err:=url.ParseQuery(r.URL.RawQuery)
	//essayid:=queryForm["sreach"][0]
	//fmt.Println("tst")
	blogtag, err := blogModel.GetBlogEssayTime()
	if err != nil {
		msg := "blogModel GetBlogEssayTime run fail:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetBlogEssayTime运行错误").Response()
		return
	}
	msg := "GetBlogTime success"
	//logger.Logger.Info(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("博客时间归档获取成功").FbData(blogtag).Response()
}

func UpdateBlogEssay(w http.ResponseWriter, r *http.Request) {
	fb := feedback.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "ReadAll failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body获取错误").Response()
		return
	}
	var essayinfo constant.BlogEssayInfo
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
	//id:=essayinfo.Id
	//new:=id.Trim(essayinfo.Id,"\"")
	//fmt.Println(essayinfo.Id)
	id, err := strconv.Atoi(essayinfo.Id)
	if err != nil {
		msg := "essayinfo.Id to int failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("essayinfo.Id转换错误").Response()
		//data:=model.FeedBackErrorHandle(501,msg)
		//fmt.Fprintln(w,string(data))
		return
	}
	blogwrite, err := blogModel.UpdateBlogEssay(essayinfo.Title, essayinfo.Author, essayinfo.Content, essayinfo.Time, essayinfo.Tag, id)
	if err != nil {
		msg := "blogModel UpdateBlogEssay run fail:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("UpdateBlogEssay运行错误").Response()
		return
	}
	if !blogwrite {
		msg := "UpdateBlog success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("该博客修改失败").Response()
		return
	}
	msg := "UpdateBlog success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该博客修改成功").Response()
}

func DeleteBlogEssay(w http.ResponseWriter, r *http.Request) {
	fb := feedback.NewFeedBack(w)
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
	essayid, ok := essayinfo["essayid"].(string)
	if !ok {
		msg := "get map key failed:" + err.Error()
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("发送body中不存在key值").Response()
		return
	}
	bloginfo, err := blogModel.DeleteBlogEssay(essayid)
	if err != nil {
		msg := "blogModel DeleteBlogEssay run fail:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("DeleteBlogEssay运行错误").Response()
		return
	}

	if !bloginfo {
		msg := "DeleteBlog success"
		//log.Println(msg)
		//logger.Logger.Info(msg)
		logger.Info(msg)
		fb.FbCode(constant.EVENT_NOT_FOUND).FbMsg("该博客删除失败").Response()
		return
	}
	msg := "DeleteBlog success"
	//log.Println(msg)
	//logger.Logger.Info(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("该博客删除成功").Response()
}
