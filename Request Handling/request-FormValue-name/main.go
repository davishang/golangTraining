// Davis Hang

// Create a webpage that serves a form and allows the user to enter their name.
// Once a user has entered their name, show their name on the webpage.
// Use req.FormValue to do this

package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)

type Name struct {
	YourName string
}

func main(){
	// parse template
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter,req *http.Request){
		// receive form submission
		name := req.FormValue("name")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Printf("%T\n", name)

		// execute template
		err = tpl.Execute(res, Name{name})
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Fatalln(err)
		}
	})
	log.Println("Listening...")
	http.ListenAndServe(":8080",nil)
}


