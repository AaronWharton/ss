package middleware

import (
	"log"
	"net/http"
	"ss/model"
	"ss/session"
)

// middleAuth implements login control, to verify whether user logined.
func MiddleAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := session.GetSessionUser(r)
		log.Println("username: ", username)
		if username != "" {
			log.Println("Last seen: ", username)
			_ = model.UpdateLastSeen(username)
		}
		if err != nil {
			log.Println("middle gets session error and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}
