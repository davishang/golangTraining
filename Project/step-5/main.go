// Davis Hang

// PROJECT STEP 5 - continuing to build our application, integrate HMAC into
// our application to ensure that nobody tampers with the cookie.

package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourKey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

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

	getCode("test")
	yourName := req.FormValue("name")
	yourAge := req.FormValue("age")

	currentUser := User{
		Name: yourName,
		Age: yourAge,
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
			Value: id.String() + " " + yourName + " " + yourAge + " " + jsonB64,
			//Secure: true
			HttpOnly: true,
		}

		cookie.Value = getCode(cookie.Value) + "|" + cookie.Value
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

