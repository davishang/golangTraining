// Davis Hang

// PROJECT STEP 2 - have the application write a cookie called "session-fino"
// with a UUID. The cookie should serve HttpOnly and you should have the
// "Secure" flag set also though comment the "Secure" flag out as we're not using https.

package main

import (
	"log"
	"strings"
	"net/http"
	"html/template"
	"github.com/nu7hatch/gouuid"
)

func foo(res http.ResponseWriter, req *http.Request) {
	// parse template
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)

//	if req.URL.Path != "/" {
//		http.NotFound(res, req)
//		return
//	}

	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
	}

	if req.FormValue("name") != "" && !strings.Contains(cookie.Value, "name") {
		cookie.Value = cookie.Value + ` name=` + req.FormValue("name")
	}

	http.SetCookie(res, cookie)
}

func main(){
	http.HandleFunc("/", foo)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}


// NOT GOOD PRACTICE
// adding user data to a cookie
// with no way of knowing whether or not
// they might have altered that data
//
// HMAC would allow us to determine
// whether or not the data in the cookie was altered
//
// however, best to store user data
// on the server
// and keep backups