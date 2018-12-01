package controller

import (
	"html/template"
	"log"
	"net/http"
	"ss/view"
)

var (
	homeController home
	templates      map[string]*template.Template
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	tpName := "login.html"
	lvm := view.LoginViewModel{}
	v := lvm.GetVM()
	if request.Method == http.MethodGet {
		_ = templates[tpName].Execute(writer, &v)
	}
	if request.Method == http.MethodPost {
		request.ParseForm()
		username := request.Form.Get("username")
		password := request.Form.Get("password")

		if len(username) < 3 {
			v.Errors("username must longer than 3")
		}
		if len(password) < 8 {
			v.Errors("password must longer than 8")
		}

		if len(v.Error) > 0 {
			_ = templates["index.html"].Execute(writer, &v)
		} else {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
		}
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	ivm := view.IVM{}
	v := ivm.GetVM()
	err := templates["index.html"].Execute(writer, &v)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	templates = populateTemplates()
}

func Startup() {
	homeController.registerRoutes()
}
