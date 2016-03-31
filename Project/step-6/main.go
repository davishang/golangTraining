// Davis Hang

// PROJECT STEP 6 - create a page which illustrates what happens when a user
// changes a cookie. You can hard-code a changed cookie into your code.

package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"crypto/hmac"
	"crypto/sha256"
	"strings"
)

type User struct {
	Name string
	Age  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template.gohtml"))
}

func foo() string {
	u := User{
		Name: "Nintendo",
		Age: "64",
	}
	bs, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return string(bs)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourKey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func handler(res http.ResponseWriter, req *http.Request) {

	data:= foo()
	code := getCode(data)
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String() + "|" + data + "|" + code,
			//Secure: true
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Fprintln(res, cookie.Value)
	xs := strings.Split(cookie.Value, "|")
	// usrToken := xs[0]
	userName := xs[1] + "say what?"
	userCode := xs[2]
	fmt.Fprintln(res, data)
	fmt.Fprintln(res, userName)
	if userCode == getCode(userName) {
		fmt.Fprintln(res, "Code valid")
	} else {
		fmt.Fprintln(res, "Code Invalid")
	}

	var u User
	err = json.Unmarshal([]byte(userName), &u)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}

	fmt.Fprintln(res, u)
	tpl.ExecuteTemplate(res, "template.gohtml", u)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Listening.....")
	http.ListenAndServe(":8080", nil)
}

