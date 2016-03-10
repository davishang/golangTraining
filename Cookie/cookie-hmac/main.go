// Create a webpage which writes a cookie to the client's machine.
// Though this is NOT A BEST PRACTICE, you will store some session
// data in the cookie. Make sure you use HMAC to ensure that session
// data is not changed by a user.

/*
Package hmac implements the Keyed-Hash Message Authentication Code (HMAC)
as defined in U.S. Federal Information Processing Standards Publication 198.
An HMAC is a cryptographic hash that uses a key to sign a message. The receiver
verifies the hash by recomputing it using the same key.

Receivers should be careful to use Equal to compare MACs
in order to avoid timing side-channels:

Package sha256 implements the SHA224 and SHA256 hash algorithms
as defined in FIPS 180-4.
*/

package main

import (
	"io"
	"fmt"
	"log"
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func handler(res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	cookie, err := req.Cookie("session-id")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: "",
			// Secure: true,
			HttpOnly: true,
		}
	}

	if req.FormValue("email") != "" {
		email := req.FormValue("email")
		cookie.Value = email + `|` + getCode(email)
	}

	fmt.Println(getCode(cookie.Value))
	http.SetCookie(res, cookie)
	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	    `+cookie.Value+`
	      <input type="email" name="email">
	      <input type="password" name-"password">
	      <input type="submit">
	    </form>
	  </body>
	</html>`)
}

func main(){
	http.HandleFunc("/", handler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
