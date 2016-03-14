// Davis Hang

// PROJECT STEP 3 - continuing to build our application, create a template
// which is a form. The form should gather the user's name and age.
// Store the user's name and age in the cookie.

package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/nu7hatch/gouuid"
)

type User struct {
	Name string
	Age string
}

func foo(res http.ResponseWriter, req *http.Request) {
	// parse template
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// receive form submission
	name := req.FormValue("name")
	age := req.FormValue("age")
	// output to console
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)

	// execute template
	tpl.Execute(res, User{name, age} )

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
		http.SetCookie(res, cookie)
	}
	// looks like this prints to console...debug later
	fmt.Println(cookie)

	//if req.FormValue("name") != "" && !strings.Contains(cookie.Value, "name") {
	//	cookie.Value = cookie.Value + ` name=` + req.FormValue("name")
	//}
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