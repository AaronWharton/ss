package controller

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"ss/utils"
)

var (
	homeController home
	templates      map[string]*template.Template

	// sessions
	SessionName string
	Store       *sessions.CookieStore
)

// init template files
func init() {
	templates = utils.PopulateTemplates()

	Store = sessions.NewCookieStore([]byte("something_secret"))
	SessionName = "ss"		// use project name to sessionName
}

// activate routes
func Start() {
	homeController.RegisterRoutes()
}

func GetUserSession(r *http.Request) (string, error) {
	var username string
	session, err := Store.Get(r, SessionName)
	if err != nil {
		return "", err
	}
	val := session.Values["user"]
	fmt.Println("value:", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get user session")
	}
	fmt.Println("username:", username)
	return username, nil
}

func SetUserSession(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := Store.Get(r, SessionName)
	if err != nil {
		return err
	}

	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := Store.Get(r, SessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}