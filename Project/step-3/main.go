// Davis Hang

// PROJECT STEP 3 - continuing to build our application, create a template
// which is a form. The form should gather the user's name and age.
// Store the user's name and age in the cookie.

package main

import ("log"
	"net/http"
	"html/template"
	"github.com/nu7hatch/gouuid"
)

type User struct {
	Name string
	Age  string
}

func handler(res http.ResponseWriter, req *http.Request) {
	// parse template
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	uName := req.FormValue("name")
	uAge := req.FormValue("age")

	cookie, err := req.Cookie("session-fino")
	if err != nil {

		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + uName + uAge,
			//Secure: true
			HttpOnly: true,
		}

		http.SetCookie(res, cookie)
	}
	err = tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
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