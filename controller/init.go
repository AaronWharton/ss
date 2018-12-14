package controller

import (
	"github.com/gorilla/sessions"
	"html/template"
	"ss/utils"
)

var (
	routerController router
	templates      map[string]*template.Template
)

// init template files
func init() {
	templates = utils.PopulateTemplates()

	store = sessions.NewCookieStore([]byte("something_secret"))
	sessionName = "ss" // use project name to sessionName
}

// activate routes
func Start() {
	routerController.RegisterRoutes()
}

