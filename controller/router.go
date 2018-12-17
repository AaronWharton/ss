package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ss/middleware"
	"ss/session"
	"ss/view"
)

type router struct{}

func (r router) RegisterRoutes() {
	m := mux.NewRouter()
	m.HandleFunc("/", middleware.MiddleAuth(indexHandler))
	m.HandleFunc("/register", registerHandler)
	m.HandleFunc("/login", loginHandler)
	m.HandleFunc("/logout", middleware.MiddleAuth(logoutHandler))
	m.HandleFunc("/user/{username}", middleware.MiddleAuth(profileHandler))
	m.HandleFunc("/profile_edit", middleware.MiddleAuth(profileEditHandler))
	http.Handle("/", m)
}

func profileEditHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "profile_edit.html"
	username, _ := session.GetSessionUser(r)
	v := view.PEV{}.GetView(username)
	if r.Method == http.MethodGet {
		err := templates[tpName].Execute(w, &v)
		if err != nil {
			log.Println(err)
		}
	}

	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		aboutme := r.Form.Get("aboutme")
		log.Println(aboutme)
		if err := view.UpdateAboutMe(username, aboutme); err != nil {
			log.Println("Update aboutme error: ", err)
			_, _ = w.Write([]byte("Error update aboutme"))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/user/%s", username), http.StatusSeeOther)
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "profile.html"
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := session.GetSessionUser(r)
	v, err := view.PVM{}.GetView(sUser, pUser)
	if err != nil {
		msg := fmt.Sprintf("user ( %s ) doesn't exist", pUser)
		_, _ = w.Write([]byte(msg))
		return
	}
	_ = templates[tpName].Execute(w, &v)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	username, _ := session.GetSessionUser(r)
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
			_ = session.SetSessionUser(w, r, username)
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
			_ = session.SetSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = session.ClearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
