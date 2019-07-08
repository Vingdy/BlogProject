//package main
//
//import (
//	"net/http"
//	"fmt"
//	"routers"
//	"utils"
//	"log"
//	"conf"
//	"constant"
//	"db"
//)
//
//func handler(w http.ResponseWriter, r *http.Request)  {
//	fmt.Fprintf(w, "hello, https")
//}
//
//func main() {
//	http.HandleFunc("/",handler)
//	http.ListenAndServeTLS(":443","1_vingdream.cn_bundle.crt","2_vingdream.cn.key", routers.SetRouter())
//}
package main

import (
	"log"
	"db"
	"utils"
	"cache"
	"logger"
	"conf"
	"constant"
	"net/http"
	"routers"
)

//var sessionMgr *session.SessionManager = nil

func main() {
	proDir,err:= utils.GetProDir()
	if err != nil {
		log.Panicln("Get ProDir err: " + err.Error())
	}
	conf.Init(proDir,constant.BUILD_TYPE_PROD)
	db.InitDB(conf.App.DBHost,conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword,conf.App.DBName,conf.App.DBDriver)
	cache.Init()

	err=logger.InitLogger(proDir)

	defer db.Db.Close()
	//defer logger.Close()
	err=http.ListenAndServeTLS(":"+conf.App.ServerPort,"1_vingdream.cn_bundle.crt","2_vingdream.cn.key",routers.SetRouter())
	//err=http.ListenAndServe(":80",routers.SetRouter())
	log.Println(err)
	if err != nil{
		log.Println("开启"+conf.App.ServerPort+"端口服务")
	}else{
		log.Println(err)
	}
}
