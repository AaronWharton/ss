package controller

import (
	"log"
	"net/http"
	"ss/view"
)

type home struct{}

func (h home) RegisterRoutes() {
	http.HandleFunc("/", middleAuth(indexHandler))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", middleAuth(logoutHandler))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = ClearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tplName := "login.html"
	lvm := view.LVM{}
	v := lvm.GetView()
	if r.Method == http.MethodGet {
		_ = templates[tplName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		// verify username and password on server
		// NOTE: it is better to verify data on web rather than server except for password
		if len(username) == 0 {
			v.AddError("Username must not be blank!")
		}
		if len(password) < 6 {
			v.AddError("Password must be longer than 8!")
		}
		if !view.CheckLogin(username, password) {
			v.AddError("username and password not correct, please try again!")
		}
		if len(v.Errors) > 0 {
			_ = templates[tplName].Execute(w, &v)

		} else {
			if err := SetSessionUser(w, r, username); err != nil {
				log.Fatalln("can't set session for user: ", username)
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	username, _ := GetSessionUser(r)
	v := view.IVM{}.GetView(username)
	_ = templates[tpName].Execute(w, &v)
}
