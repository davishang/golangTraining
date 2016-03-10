// Davis Hang
// PROJECT STEP 1 - create a web application that serves an HTML template.

package main

import (
	"os"
	"log"
	"strings"
	"html/template"
)

type Form struct {
	Name string
}

func main() {
	var err error

	tpl := template.New("template.gohtml")
	tpl = tpl.Funcs(template.FuncMap{
		"uppercase": func(str string) string {
			return strings.ToUpper(str)
		},
	})
	tpl, err = tpl.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, Form{
		Name: "Your Page",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

