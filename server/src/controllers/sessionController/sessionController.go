package sessionController

import (
	"constant"
	"fmt"
	"log"
	"logger"
	"net/http"
	"session"
	"utils/feedback"
)

//func ActionLogger() http.HandlerFunc {
//	return middleWare.MiddleWareHandler(actionLogger, controller.Permission{
//		LevelLim: constant.USER_PLAYER,
//	})
//}

func GetRole(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	if w == nil || r == nil {
		log.Println("gate fail: w or r nil")
		return
	}
	fb := feedback.NewFeedBack(w)
	if r.Method == http.MethodOptions {
		fb.FbCode(0).Response()
		return
	}
	sess, err := session.GetSession(r)
	if err != nil {
		log.Println("gateway judge session fail: " + err.Error())
		fb.FbCode(constant.SYS_ERR).Response()
		return
	}
	fmt.Println(sess)
	if sess == nil {
		fb.FbCode(constant.SESSION_EXPIRED).Response()
		return
	}
	role := sess.Role
	msg := "GetRole Success"
	//logger.Logger.Error(msg)
	logger.Info(msg)
	//log.Println(msg)
	fb.FbCode(constant.SUCCESS).FbData(role).FbMsg("获取权限等级成功").Response()
	defer r.Body.Close()
}
