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
	m.HandleFunc("/follow/{username}", middleware.MiddleAuth(followHandler))
	m.HandleFunc("/unfollow/{username}", middleware.MiddleAuth(unfollowHandler))
	m.HandleFunc("/explore", middleware.MiddleAuth(exploreHandler))
	http.Handle("/", m)
}

func exploreHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "explore.html"
	username, _ := session.GetSessionUser(r)
	current := getPage(r)
	v := view.EVM{}.GetView(username, current, pageLimit)
	_ = templates[tpName].Execute(w, &v)
}

func followHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := session.GetSessionUser(r)

	err := view.Follow(sUser, pUser)
	if err != nil {
		log.Println("Follow error: ", err)
		_, _ = w.Write([]byte("Error in Follow()"))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusSeeOther)
}

func unfollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := session.GetSessionUser(r)

	err := view.Unfollow(sUser, pUser)
	if err != nil {
		log.Println("Unfollow error: ", err)
		_, _ = w.Write([]byte("Error in Unfollow()"))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusSeeOther)
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
		about := r.Form.Get("about")
		log.Println(about)
		if err := view.UpdateAboutMe(username, about); err != nil {
			log.Println("Update about error: ", err)
			_, _ = w.Write([]byte("Error update about"))
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
	current := getPage(r)
	v, err := view.PVM{}.GetView(sUser, pUser, current, pageLimit)
	if err != nil {
		msg := fmt.Sprintf("user ( %s ) doesn't exist", pUser)
		_, _ = w.Write([]byte(msg))
		return
	}
	_ = templates[tpName].Execute(w, &v)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	current := getPage(r)
	username, _ := session.GetSessionUser(r)
	if r.Method == http.MethodGet {
		flash := getFlash(w, r)
		v := view.IVM{}.GetView(username, flash, current, pageLimit)
		_ = templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		body := r.Form.Get("body")
		errMsg := checkLen("Post", body, 1, 180)
		if errMsg != "" {
			setFlash(w, r, errMsg)
		} else {
			err := view.CreatePost(username, body)
			if err != nil {
				log.Println("Add post error: ", err)
				_, _ = w.Write([]byte("Error insert post in database"))
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
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
