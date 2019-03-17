package rts

import (
	"net/http"

	"github.com/alecthomas/template"
	"github.com/parallelcointeam/jorm9/hlp"
	"github.com/parallelcointeam/jorm9/rps"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("").ParseFiles("./tpl/admin/login.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/main.gohtml", "./tpl/admin/nav.gohtml")
	tmpl.ExecuteTemplate(w, "base", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if !hlp.IsEmpty(name) && !hlp.IsEmpty(pass) {
		_userIsValid := rps.UserIsValid(name, pass)

		if _userIsValid {
			SetCookie(name, w)
			redirectTarget = "/admin/"
		} else {
			redirectTarget = "/"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	ClearCookie(w)
	http.Redirect(w, r, "/", 302)
}

func SetCookie(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func ClearCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func GetUserName(w *http.Request) (userName string) {
	if cookie, err := w.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
