// Davis Hang

// PROJECT STEP 4 - refactoring our application, create a new data type called
// "user" which has fields for the user's name and age. When you receive the
// user's name and age form submission, create a variable of type "user"
// then put those values from the form submission into the fields for that
// variable. Marshal your variable of type "user"  to JSON. Encode that
// JSON to base64. Store that value in the cookie.


// testing encoding/json
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

func main() {
	data := []byte(`{"age": 64, "name": "Nintendo"}`)
	p := User{}
	json.Unmarshal(data, &p)
	fmt.Println(p)
}

/*
func main() {
	p := User{"Nintendo", 64}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes))
}

// Field name tag not respected when encoding if there is a space after ":"

*/


/************************************
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

*///////////////////////////////