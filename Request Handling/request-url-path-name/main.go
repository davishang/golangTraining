// Create a webpage that serves at localhost:8080 and will display the name
// in the url when the url is localhost:8080/name - use req.URL.Path to do this

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	// The trailing [1:] means "create a sub-slice of Path from the 1st
	// character to the end." This drops the leading "/" from the path name.
	fmt.Fprint(res, "Name: ", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}