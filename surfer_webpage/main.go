package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func webpage(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "webpage ran")
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	http.ServeFile(res, req, "html/index.html")
	tpl.Execute(res, nil)
}

func main() {
	//Handle css directory.
	http.Handle("/css/",http.StripPrefix("/css",http.FileServer(http.Dir("css"))))
	//Handle pic directory.
	http.Handle("/pic/",http.StripPrefix("/pic",http.FileServer(http.Dir("pic"))))
	//Handle index.html.
	http.HandleFunc("/page/", webpage)
	http.ListenAndServe(":8080", nil)
}

