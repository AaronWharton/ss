package controller

import (
	"net/http"
	"ss/view"
)

type home struct {}

func (h home) RegisterRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
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
		if len(password) < 8 {
			v.AddError("Password must be longer than 8!")
		}
		if len(v.Errors) > 0 {
			_ = templates[tplName].Execute(w, &v)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	v := view.IVM{}.GetView()
	_ = templates["index.html"].Execute(w, &v)
}