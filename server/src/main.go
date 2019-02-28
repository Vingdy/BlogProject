package main

import (
	"net/http"
	"log"
	"routers"
	"db"
	"logger"
	"utils"
	"cache"
)

//var sessionMgr *session.SessionManager = nil

func main() {
	db.InitDB("localhost","5432", "postgres", "VingB2by","test","postgres")
	cache.Init()
	proDir,err:= utils.GetProDir()
	//fmt.Println(err)
	if err != nil {
		log.Panicln("Get ProDir err: " + err.Error())
	}
	//fmt.Println(proDir)
	err=logger.InitLogger(proDir)
	if err != nil {
		log.Panicln("InitLogger err: " + err.Error())
	}
	log.Println("====开始监听8080端口=====")
	//http.HandleFunc("/", login.Login)\
	//ExePath, err := utils.InitDirectories()
	//if err != nil {
	//	panic(err)
	//}
	//err=logger.InitLogger(ExePath)
	//if err != nil {
	//	panic(err)
	//}
	//Init()

	defer db.Db.Close()
	defer logger.Close()
	err=http.ListenAndServe(":"+"8080", routers.SetRouter())
	if err != nil{
		log.Println("开启8080端口服务")
	}

}
//package main
//
//import (
//	"fmt"
//	"time"
//	"runtime"
//	"bytes"
//	"blog/framework/constant"
//	"blog/framework/conf"
//	"blog/system/untils"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"blog/framework/logger"
//
//	"github.com/gorilla/mux"
//)
//
////build的时候注入的变量
////构建版本
//var buildVer string
//
////构建类型 开发测试/发布
//var buildType string
//
//func handlePanic() {
//	if err := recover(); err != nil {
//		formatErr := fmt.Sprintf("%s\t%v", time.Now().Local(), err)
//		stack := make([]byte, 1024*10)
//		runtime.Stack(stack, true)
//		stack = bytes.Replace(stack, []byte("\u0000"), []byte(""), -1)
//		fmt.Printf("%s\n%s", formatErr, stack)
//		ioutil.WriteFile(conf.App.ProDir+"/panic", bytes.Join([][]byte{[]byte(formatErr), stack}, []byte("\n")), 0644)
//	}
//}
//
//func setupEnv(buildType string) {
//	if buildType == "" {
//		buildType = constant.BUILD_TYPE_DEV
//	}
//	fmt.Println("Current BuildType:"+buildType)
//	proDir, err := utils.GetProDir()
//	fmt.Println("Current ProDir:" + proDir)
//	if err != nil {
//		log.Panicln("Get ProDir err: " + err.Error())
//	}
//
//	//conf.Init(proDir, buildType)
//
//	//fileController.InitFileSystem()
//
//	//db.InitDB(conf.App.DBHost, conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword, conf.App.DBName, conf.App.DBDriver)
//
//	//cache.Init()
//
//	//cacheController.Init()
//
//	//session.Init()
//
//	//err = logger.Init(conf.App.ProDir)
//	//if err != nil {
//	//	log.Panicln("Init Logger err: " + err.Error())
//	//}
//
//	//view.Init()
//
//	//if conf.App.AppEnv==constant.BUILD_TYPE_PROD{
//	//	initWX()
//	//}
//
//}
//
//func main() {
//	defer handlePanic()
//	setupEnv(buildType)
//	//defer cache.Cache.Close()
//	//defer db.Db.Close()
//	//defer logger.Sync()
//
//	fmt.Println("Server listening at port " + conf.App.ServerPort)
//
//	err := http.ListenAndServe(":"+conf.App.ServerPort,setRouter())
//	if err != nil {
//		panic(err)
//		//logger.Fatal(err)
//	}
//}
//
//func setRouter() *mux.Router {
//	//r := mux.NewRouter()
//
//	//test network api
//	//r.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
//	//	var ttt int
//	//	err := db.Db.QueryRow("SELECT 1+1 AS TTT").Scan(&ttt)
//	//	if err != nil {
//	//		fmt.Fprint(w, "ERROR")
//	//		return
//	//	}
//	//	if ttt != 2 {
//	//		fmt.Fprint(w, "NO")
//	//		return
//	//	}
//	//	fmt.Fprint(w, `ok`)
//	//}).Methods(http.MethodGet, http.MethodOptions)
//
//
//	////others
//	//r.HandleFunc("/api/layout", assistantController.Layout).Methods(http.MethodOptions, http.MethodGet)
//	//r.HandleFunc("/api/frontlog", frontLogController.ReceiveFrontLog).Methods(http.MethodOptions, http.MethodPost)
//	//
//	//if true { //buildType == constant.BUILD_TYPE_DEV {
//	//	r.HandleFunc("/api/mocklogin", mockLoginController.MockLogin).Methods(http.MethodOptions, http.MethodGet)
//	//}
//
//	//r.HandleFunc("/api/tmp", func(w http.ResponseWriter, r *http.Request) {
//	//	statisticsController.TimmingStats()
//	//	fmt.Fprint(w, "ok")
//	//}).Methods(http.MethodOptions, http.MethodGet)
//	//
//	////TODO 权限控制加细
//	////TODO 其实应该为 管理员+参赛者+游客(微信内)
//	//fileServerHandler := http.StripPrefix(conf.App.FileDownloadPrefix, http.FileServer(http.Dir(conf.App.FileDir)))
//	//r.PathPrefix(conf.App.FileDownloadPrefix).Handler(controller.GateWay(fileServerHandler.ServeHTTP, controller.Permission{
//	//	constant.USER_ALL_PERMIT,
//	//}))
//	//
//	//err := buildUtils.HandleInjectVariable(buildVer, buildType, &conf.App.WebPageHost, r)
//	//if err != nil {
//	//	logger.Error("handleInjectVariable failed " + err.Error())
//	//}
//	//
//	//setWebPageHostHandle, err := buildUtils.SetWebPageHost(&conf.App.WebPageHost)
//	//if err != nil {
//	//	logger.Error("SetWebPageHost failed " + err.Error())
//	//} else {
//	//	if setWebPageHostHandle != nil {
//	//		r.HandleFunc("/set/wph", setWebPageHostHandle).Methods(http.MethodOptions, http.MethodGet)
//	//	}
//	//}
//	//
//	//initFrontEndMux(r)
//	//
//	//if buildType == constant.BUILD_TYPE_DEV {
//	//	r.PathPrefix("/schoolnoc").Handler(http.StripPrefix("/schoolnoc", http.FileServer(http.Dir(conf.App.TemplateDir))))
//	//}
//	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(conf.App.TemplateDir))))
//
//	return r
//}
