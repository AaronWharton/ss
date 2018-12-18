package controller

import (
	"github.com/gorilla/sessions"
	"html/template"
	"ss/session"
	"ss/utils"
)

var (
	routerController router
	templates        map[string]*template.Template
	flashName            string
)

// init template files
func init() {
	templates = utils.PopulateTemplates()

	session.Store = sessions.NewCookieStore([]byte("something_secret"))
	session.SessionsName = "ss" // use project name to sessionName

	flashName = "ss-flash"
}

// activate routes
func Start() {
	routerController.RegisterRoutes()
}
