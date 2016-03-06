// Davis Hang

// Create a webpage that serves a form and allows the user to upload a txt file.
// You do not need to check if the file is a txt; bad programming but just
// trust the user to follow the instructions. Once a user has uploaded a txt
// file, copy the text from the file and display it on the webpage.
// Use req.FormFile and io.Copy to do this

package main

import (
	"io"
	"os"
	"log"
	"net/http"
	"html/template"
	"path/filepath"
)

func main(){
	// parse template
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// handler
	http.HandleFunc("/", func(res http.ResponseWriter,req *http.Request) {
		// receive form submission
		// POST takes the uploaded file(s) and saves it to disk.
		if req.Method == "POST" {
			src, _, err := req.FormFile("file")
			if err != nil {
				panic(err)
			}
			defer src.Close()
			// create destination file
			dst, err := os.Create(filepath.Join("./", "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dst.Close()
			// copy the uploaded file to the destination file
			io.Copy(dst, src)
		}

		// execute template
		err = tpl.Execute(res, nil)
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Println(err)
		}
	})
	log.Println("Listening...")
	http.ListenAndServe(":8080",nil)
}

