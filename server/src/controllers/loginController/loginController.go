package loginController

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"utils/feedback"
	"models/loginModel"
	"constant"
	"session"
	"strconv"
	"logger"
)

//func ActionLogger() http.HandlerFunc {
//	return middleWare.MiddleWareHandler(actionLogger, controller.Permission{
//		LevelLim: constant.USER_PLAYER,
//	})
//}

func Login(w http.ResponseWriter, r *http.Request) {

	//var fb feedback.FeedBack
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

	var logininfo constant.LoginInfo
	err = json.Unmarshal(body, &logininfo)
	if err != nil {
		msg := "json Unmarshal failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("请求body解析json错误").Response()
		//fb.Response(w, constant.PARA_ERR, "请求body解析json错误", nil)
		return
	}
	//fmt.Println(&logininfo)
	//username:=logininfo.LoginAccount
	//password:=logininfo.LoginPassword
	//role:=logininfo.Role
	//fmt.Println(logininfo.LoginAccount,logininfo.LoginPassword,logininfo.Role)
	if len(logininfo.LoginAccount)<0{
		msg:="LoginAccount is empty"
		//logger.Logger.Error(msg)
		//log.Println(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("LoginAccount参数为空").Response()
		return
	}
	if len(logininfo.LoginPassword)<0{
		msg:="LoginPassword is empty"
		//logger.Logger.Error(msg)
		//log.Println(msg)
		logger.Info(msg)
		fb.FbCode(constant.PARA_ERR).FbMsg("LoginPassword参数为空").Response()
		return
	}
	//if len(logininfo.Role)<0{
	//	msg:="Role is empty"
	//	logger.Logger.Error(msg)
	//	fb.FbCode(constant.PARA_ERR).FbMsg("Role参数为空").Response()
	//	return
	//}
	//Todo 不合法参数检查
	isexist, err := loginModel.CheckLoginAcc(logininfo.LoginAccount)
	if err != nil {
		msg := "LoginModel run failed:" + err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SYS_ERR).FbMsg("CheckLoginAcc运行错误").Response()
		//fb.Response(w, constant.SYS_ERR, "", nil)
		return
	}
	if !isexist {
		msg := "LoginAccount is not exist"
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.ADMIN_NOT_EXIST).FbMsg("账号不存在").Response()
		//fb.Response(w, constant.ADMIN_NOT_EXIST, "账号不存在", nil)
		return
	}
	loginok,err := loginModel.Login(logininfo.LoginAccount, logininfo.LoginPassword)
	//fmt.Println(err)
	if err != nil {
		msg := "LoginModel run failed:"+err.Error()
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.DB_ERR).FbMsg("Login运行错误").Response()
		//fb.Response(w, constant.DB_ERR, msg, nil)
		return
	}
	if len(loginok)==0{
		msg := "LoginAccount is not exist"
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.ADMIN_PWD_WRONG).FbMsg("密码错误").Response()
		//fb.Response(w, constant.ADMIN_NOT_EXIST, "账号不存在", nil)
		return
	}
//if loginok[0].LoginPassword!=logininfo.LoginPassword{
//		msg := "LoginPassword is wrong"
//		log.Println(msg)
//		fb.FbCode(constant.ADMIN_PWD_WRONG).FbMsg("密码错误").Response()
//		//fb.Response(w, constant.ADMIN_NOT_EXIST, "账号不存在", nil)
//		return
//	}
	msg := "Login success"
	//log.Println(msg)
	//logger.Info(msg)
	role_int, err := strconv.Atoi(loginok[0].Role)
	err = session.SetSession(w, r, &constant.Session{
		 loginok[0].LoginAccount, role_int,
	})
	//sess,_:=session.GetSession(r)
	//fmt.Println(sess)
	//logger.Logger.Error(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("登陆成功").Response()
	defer r.Body.Close()
}

func LogOut(w http.ResponseWriter,r *http.Request){
	fb:=feedback.NewFeedBack(w)
	err:=session.DestroySession(r)
	if err!=nil{
		msg:="LogOut DestroySession Failed"
		//log.Println(msg)
		//logger.Logger.Error(msg)
		logger.Info(msg)
		fb.FbCode(constant.SESSION_EXPIRED).FbMsg("LogOut运行失败").Response()
	}
	msg := "LogOut success"
	//log.Println(msg)

	//logger.Logger.Error(msg)
	//log.Println(msg)
	logger.Info(msg)
	fb.FbCode(constant.SUCCESS).FbMsg("退出成功").Response()
}

