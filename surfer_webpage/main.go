package main

import (
	"html/template"
	//"io"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {

	//io.WriteString(res, "webpage ran")
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	//http.ServeFile(res, req, "index.html")
	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/css/",http.StripPrefix("/css",http.FileServer(http.Dir("css") ) ) )
	http.Handle("/pic/",http.StripPrefix("/pic",http.FileServer(http.Dir("pic") ) ) )
	http.ListenAndServe(":8080", nil)
}

// localhost:8080
