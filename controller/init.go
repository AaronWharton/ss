package controller

import (
	"html/template"
	"ss/utils"
)

var(
	homeController home
	templates	map[string]*template.Template
)

// init template files
func init() {
	templates = utils.PopulateTemplates()
}

// activate routes
func Start() {
	homeController.RegisterRoutes()
}
