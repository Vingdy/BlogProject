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
	"cache"
	"conf"
	"constant"
	"db"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"log"
	"logger"
	"net/http"
	"routers"
	"utils"
)

//var sessionMgr *session.SessionManager = nil

func main() {
	proDir, err := utils.GetProDir()
	if err != nil {
		log.Panicln("Get ProDir err: " + err.Error())
	}
	conf.Init(proDir, constant.BUILD_TYPE_PROD)
	db.InitDB(conf.App.DBHost, conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword, conf.App.DBName, conf.App.DBDriver)
	cache.Init()

	err = logger.InitLogger(proDir)

	defer db.Db.Close()
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	//defer logger.Close()
	err = http.ListenAndServeTLS(":"+conf.App.ServerPort, "1_vingdream.cn_bundle.crt", "2_vingdream.cn.key", routers.SetRouter())

	//err=http.ListenAndServe(":80",routers.SetRouter())
	if err == nil {
		log.Println(err)
	}
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "vingdream.cn:443",
		})

		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}

		c.Next()
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	/*_host := strings.Split(req.Host, ":")
	_host[1] = "443"

	target := "https://" + strings.Join(_host, ":") + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}*/

	http.Redirect(w, r, "https://vingdream.cn"+r.RequestURI, http.StatusTemporaryRedirect)
}
