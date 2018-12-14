package controller

import (
	"log"
	"net/http"
)

// middleAuth implements login control, to verify whether user logined.
func middleAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username ,err := GetSessionUser(r)
		log.Println("username: ", username)
		if err != nil {
			log.Println("middle gets session error and redirect to login: ", err)
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}
