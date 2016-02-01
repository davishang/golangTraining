package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Foo ran")
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	http.ServeFile(res, req, "index.html")
	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/page/", foo)
	http.ListenAndServe(":8080", nil)
}

