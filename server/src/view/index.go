package view

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"log"
	"constant"
)

var (
	indexTemplate *template.Template
	templates     *template.Template
)

func Init() {
	var allfile []string
	files, err := ioutil.ReadDir(conf.App.TemplateDir)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".html") {
			allfile = append(allfile, conf.App.TemplateDir+fileName)
		}
	}
	templates = template.Must(template.ParseFiles(allfile...))
	indexTemplate = templates.Lookup("index.html")
	log.Println("view init finish")
}

func LoadTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.Method == "GET" {
		indexTemplate.Execute(w, nil)
		return
	}
	http.Redirect(w, r, constant.RE_PROXY_SCHEME, http.StatusFound)
}
