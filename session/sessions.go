package session

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var SessionsName string

var Store *sessions.CookieStore

func GetSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := Store.Get(r, SessionsName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	fmt.Println("value:", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}

	fmt.Println("username:", username)
	return username, nil
}

func SetSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := Store.Get(r, SessionsName)
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
	session, err := Store.Get(r, SessionsName)
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
