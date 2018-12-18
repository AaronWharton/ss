package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"ss/session"
	"ss/view"
)

// Login Check
func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}

func checkUsername(username string) string {
	return checkLen("Username", username, 3, 20)
}

func checkPassword(password string) string {
	return checkLen("Password", password, 6, 20)
}

func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}

func checkUserPassword(username, password string) string {
	if !view.CheckLogin(username, password) {
		return fmt.Sprintf("Username and password is not correct.")
	}
	return ""
}

func checkUserExist(username string) string {
	if !view.CheckUserExist(username) {
		return fmt.Sprintf("Username already exists, please choose another username")
	}
	return ""
}

func checkLogin(username, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserPassword(username, password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

func checkRegister(username, email, pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "2 password does not match")
	}
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserExist(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// addUser()
func addUser(username, password, email string) error {
	return view.AddUser(username, password, email)
}

// flash message
func setFlash(w http.ResponseWriter, r *http.Request, message string) {
	session2, _ := session.Store.Get(r, session.SessionsName)
	session2.AddFlash(message, flashName)
	_ = session2.Save(r, w)
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	session2, _ := session.Store.Get(r, session.SessionsName)
	fm := session2.Flashes(flashName)
	if fm == nil {
		return ""
	}

	_ = session2.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}
