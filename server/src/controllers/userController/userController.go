package userController

import (
	"constant"
	"encoding/json"
	"io/ioutil"
	"logger"
	"models/userModel"
	"net/http"
	"utils/feedback"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {
	fb := feedback.NewFeedBack(w)
	userdata, err := userModel.GetUserData()
	if err != nil {
		msg := "userModel GetUserData run fail:" + err.Error()
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetUserData运行错误").Response()
		return
	}
	msg := "GetUserData success"
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("个人资料获取成功").FbData(userdata).Response()
}

func UpdateUserData(w http.ResponseWriter, r *http.Request) {
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
	var userdatainfo constant.UserInfo
	err = json.Unmarshal(body, &userdatainfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		//fb.Response(w, constant.PARA_ERR, "请求body解析json错误", nil)
		return
	}
	userdataok, err := userModel.UpdateUserData(userdatainfo.Name, userdatainfo.Headpicture, userdatainfo.Info)
	if err != nil {
		msg := "userModel UpdateUserData run fail:" + err.Error()
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("GetUserData运行错误").Response()
		return
	}
	msg := "UpdateUserData success"
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("个人资料更新成功").FbData(userdataok).Response()
}
