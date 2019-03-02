package routers

import (
	"github.com/gorilla/mux"
	"controllers/loginController"
	"fmt"
	"time"
	"runtime"
	"bytes"
	"net/http"
	"log"
	"utils/feedback"
	"session"
	"constant"
	"controllers/blogController"
	"controllers/gameController"
	"controllers/sentenceController"
	"controllers/imgController"
	"controllers/sessionController"
	"controllers/drawController"
	//"utils"
	"utils"
)

var path,err=utils.GetProDir()

//多路复用器根据请求返回特定的处理器

func SetRouter() *mux.Router {
	Router:=mux.NewRouter()

	if err!=nil{
		return Router
	}

	//path+=`\files`
	//fmt.Println(path)
	filesever:=http.StripPrefix("/static/", http.FileServer(http.Dir(path+`\files`)))
	Router.PathPrefix("/static/").HandlerFunc(filesever.ServeHTTP)


	Router.PathPrefix("/static").Handler(http.StripPrefix("/static",http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		//http.FileServer(http.Dir("F:\\CST\\会议视频系统\\Videoconference\\html\\dist")).ServeHTTP(w,r)
		http.FileServer(http.Dir(path+`/dist/htmlpage`)).ServeHTTP(w,r)
		//http.FileServer(http.Dir("/root/videoconference/html")).ServeHTTP(w,r)
	})))
	Router.PathPrefix("/assets").Handler(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		//http.FileServer(http.Dir("F:\\CST\\会议视频系统\\Videoconference\\html\\dist")).ServeHTTP(w,r)
		http.FileServer(http.Dir(path+`/dist/htmlpage`)).ServeHTTP(w,r)
		//http.FileServer(http.Dir("/root/videoconference/html")).ServeHTTP(w,r)
	}))
	Router.NotFoundHandler =  http.HandlerFunc(notFound)
	//Router.PathPrefix("/assets").Handler(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
	//	//http.FileServer(http.Dir("F:\\CST\\会议视频系统\\Videoconference\\html\\dist")).ServeHTTP(w,r)
	//	http.FileServer(http.Dir("D:\\SoloWork\\web-show\\blog\\htmlpage\\dist")).ServeHTTP(w,r)
	//	//http.FileServer(http.Dir("/root/videoconference/html")).ServeHTTP(w,r)
	//}))


	Router.HandleFunc("/api/login", AllowOrigin(loginController.Login)).Methods("POST","OPTIONS")
	Router.HandleFunc("/api/logout", AllowOrigin(loginController.LogOut)).Methods("GET","OPTIONS")

	Router.HandleFunc("/api/writeblogessay", AllowOrigin(GateWay(blogController.WriteBlogEssay))).Methods("POST","OPTIONS")
	Router.HandleFunc("/api/getallblogessay", AllowOrigin(blogController.GetAllBlogEssay)).Methods("GET","OPTIONS")
	Router.HandleFunc("/api/getoneblogessay", AllowOrigin(blogController.GetOneBlogEssay)).Methods("GET","OPTIONS")

	Router.HandleFunc("/api/writegameessay", AllowOrigin(gameController.WriteGameEssay)).Methods("POST","OPTIONS")
	Router.HandleFunc("/api/getallgameessay", AllowOrigin(gameController.GetAllGameEssay)).Methods("GET","OPTIONS")
	Router.HandleFunc("/api/getonegameessay", AllowOrigin(gameController.GetOneGameEssay)).Methods("GET","OPTIONS")

	Router.HandleFunc("/api/writesentence", AllowOrigin(sentenceController.WriteSentence)).Methods("POST","OPTIONS")
	Router.HandleFunc("/api/getallsentence", AllowOrigin(sentenceController.GetAllSentence)).Methods("GET","OPTIONS")

	Router.HandleFunc("/api/writedrawpicture", AllowOrigin(drawController.WriteDrawPicture)).Methods("POST","OPTIONS")
	Router.HandleFunc("/api/getalldrawpicture", AllowOrigin(drawController.GetAllDrawPicture)).Methods("GET","OPTIONS")
	Router.HandleFunc("/api/getonedrawpicture", AllowOrigin(drawController.GetOneDrawPicture)).Methods("GET","OPTIONS")

	Router.HandleFunc("/api/uploadimage",AllowOrigin(imgController.UploadPic)).Methods("POST","OPTIONS")

	Router.HandleFunc("/api/getrole",AllowOrigin(sessionController.GetRole)).Methods("GET","OPTIONS")
	//Router.HandleFunc("/api/getoneblogessay", AllowOrigin(blogController.GetOneBlogEssay)).Methods("GET")
	//r.HandleFunc("/api/get", loginController.Get).Methods("POST")

	return Router
}

func MiddleWareHandler(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
//***
//TODO http.HandlerFunc和http.Handler的区别
//***
//func TESTHandler(next http.HandlerFunc)http.HandlerFunc{
//	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
//		fmt.Println("test")
//		next.ServeHTTP(w, r)
//	})
//}
func notFound(w http.ResponseWriter, r *http.Request){
	//http.ServeFile(w,r,"F:\\CST\\会议视频系统\\Videoconference\\html\\dist\\index.html")
	http.ServeFile(w,r,path+`/dist/htmlpage`)
	//http.ServeFile(w,r,"/root/videoconference/html")
}

func GateWay(next http.HandlerFunc)http.HandlerFunc {
	//fmt.Println("before return")//这里只运行一次
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		//Todo before mux
		if r.Method=="OPTIONS"{
			return
		}
		start:=time.Now()
		defer handlePanic()
		if w==nil || r==nil{
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
		if sess == nil {
			fb.FbCode(constant.SESSION_EXPIRED).Response()
			return
		}
		fmt.Println(sess)
		//level := sess.Role
		if sess.Role>5{
			next(w,r)
		}else{
			fb.FbCode(constant.NO_AUTH).Response()
			return
		}
		end:=time.Now()
		endtime:=end.Sub(start)
		//Todo after mux
		fmt.Println(endtime)
	})
}

func AllowOrigin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		next(w, r)
	})
}

func handlePanic() {
	if err := recover(); err != nil {
		formatErr := fmt.Sprintf("%s\t%s", time.Now().Local(), err)
		stack := make([]byte, 1024*4)
		runtime.Stack(stack, false)
		stack = bytes.Replace(stack, []byte("\u0000"), []byte(""), -1)
		fmt.Println(formatErr)
		//ioutil.WriteFile(conf.App.ProDir+"/panic", bytes.Join([][]byte{[]byte(formatErr), stack}, []byte("\n")), 0644)
		//logger.Error(fmt.Sprintf("%s\n%s", err, stack))
	}
}

