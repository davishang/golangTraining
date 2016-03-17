// Davis Hang

// PROJECT STEP 4 - refactoring our application, create a new data type called
// "user" which has fields for the user's name and age. When you receive the
// user's name and age form submission, create a variable of type "user"
// then put those values from the form submission into the fields for that
// variable. Marshal your variable of type "user"  to JSON. Encode that
// JSON to base64. Store that value in the cookie.

package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
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

	currentUser := User{
		Name: uName,
		Age: uAge,
	}

	bs, err := json.Marshal(currentUser)
	if err != nil{
		fmt.Println(err)
	}

	jsonB64 := base64.StdEncoding.EncodeToString(bs)

	cookie, err := req.Cookie("session-fino")
	if err != nil {

		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + " " + uName + " " + uAge + " " + jsonB64,
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

/* testing encoding/json
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
*/
