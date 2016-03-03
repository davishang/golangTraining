//Create a webpage that displays the URL path using req.URL.Path

package main

import (
	"io"
	"log"
	"net/http"
)

func display_url_path(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, req.URL.Path)
	//res.Header().Set("Content-Type", "text/html; charset=utf-8")
	//io.WriteString(res, `<h1> `+req.URL.Path+` </h1><br>`)
}

func main() {
	http.HandleFunc("/", display_url_path)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

/*
package main

import (
	"io"
	"net/http"
)

type Snoop int

func (h Snoop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
res.Header().Set("Content-Type", "text/html; charset=utf-8")
io.WriteString(res, `<h1>`+req.URL.Path+`</h1><br>`)
}

func main() {
var dog Snoop

mux := http.NewServeMux()
mux.Handle("/", dog)

http.ListenAndServe(":8080", mux)
}
*/