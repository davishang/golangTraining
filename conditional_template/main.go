package main

import (
	"html/template"
	"log"
	"os"
)

type Person struct {
	Name string
	First string
	Last string
	Input string
}

func main() {

	home := Person{
		Name: "Student Learning Golang",
		First: "David",
		Last: "Hang",
		Input: `<script>alert("Error");</script>`,
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}
// test