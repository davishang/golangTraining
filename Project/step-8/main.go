// PROJECT STEP 8 - Allow the user to logout. Show a log-in button when the
// user is not logged-in. Show a log-out button only when the user is logged in.

package main

import (
	"log"
	"strings"
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("template/*.gohtml")
}

func index(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)

	if req.Method == "POST" {
		src, hdr, err := req.FormFile("data")
		if err != nil {
			log.Println("error uploading photo: ", err)
			// TODO: create error page to show user
		}
		cookie = uploadPhoto(src, hdr, cookie)
		http.SetCookie(res, cookie)
	}

	m := Model(cookie)
	tpl.ExecuteTemplate(res, "index.gohtml", m)
}

func logout(res http.ResponseWriter, req *http.Request) {
	cookie := newVisitor()
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", 302)
	tpl.ExecuteTemplate(res, "logout.gohtml", nil)
}

func login(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)

	if req.Method == "POST" && req.FormValue("password") == "secret" {
		m := Model(cookie)
		m.State = true
		m.Name = req.FormValue("name")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie := currentVisitor(m, id)
		http.SetCookie(res, cookie)

		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func genCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session-fino")
	if err != nil {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	// make sure set cookie uses our current structure
	if strings.Count(cookie.Value, "|") != 2 {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if tampered(cookie.Value) {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	return cookie
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fs := http.FileServer(http.Dir("photos"))
	http.Handle("/imgs/", fs)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
