//Create a webpage which uses a cookie to track the number of visits of a user.
//Display the number of visits. Make sure that the favicon.ico requests are
//not also incrementing the number of visits.
package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main(){
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// disregard favicon requests
		if req.URL.Path != "/" {
			http.NotFound(res, req)
			return
		}
		// get the cookie
		cookie, err := req.Cookie("my-cookie")
		// create cookie if one doesn't exist
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "my-cookie",
				Value: "0",
			}
		}
		// there is a cookie
		count, _ := strconv.Atoi(cookie.Value)
		count++
		cookie.Value = strconv.Itoa(count)

		http.SetCookie(res, cookie)

		io.WriteString(res, cookie.Value)
	})
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}