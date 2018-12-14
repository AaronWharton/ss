package controller

import (
	"log"
	"net/http"
	"ss/view"
)

type router struct{}

func (r router) RegisterRoutes() {
	http.HandleFunc("/", middleAuth(indexHandler))
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", middleAuth(logoutHandler))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	username, _ := GetSessionUser(r)
	v := view.IVM{}.GetView(username)
	_ = templates[tpName].Execute(w, &v)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tplName := "register.html"
	v := view.RVM{}.GetView()

	if r.Method == http.MethodGet {
		_ = templates[tplName].Execute(w, &v)
	} else if r.Method == http.MethodPost {
		_ = r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errors := checkRegister(username, email, pwd1, pwd2)
		v.AddError(errors...)
		if len(v.Errors) > 0 {
			_ = templates[tplName].Execute(w, &v)
		} else {
			if err := addUser(username, pwd1, email); err != nil {
				log.Println("Failed to add user: ", err)
				_, _ = w.Write([]byte("Error insert to database"))
				return
			}
			_ = SetSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tplName := "login.html"
	v := view.LVM{}.GetView()
	if r.Method == http.MethodGet {
		_ = templates[tplName].Execute(w, &v)
	} else if r.Method == http.MethodPost {
		_ = r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		// verify username and password on server
		// NOTE: it is better to verify data on web rather than server except for password
		errors := checkLogin(username, password)
		v.AddError(errors...)
		if len(v.Errors) > 0 {
			_ = templates[tplName].Execute(w, &v)
		} else {
			_ = SetSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = ClearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
