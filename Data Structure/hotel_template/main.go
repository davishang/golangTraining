package main

import (
	"text/template"
	"log"
	"os"
)
//Create a data structure to pass to a template which contains
//information about California hotels including Name, Address, City, Zip, Region
//Region can be: Southern, Central, Norhtern
//can hold an unlimited number of hotels
func main(){

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := region{
		Region: "Northern",
		Hotels: []hotel{
			hotel{
				Name: "North Hotel",
				Address: "123 Street",
				City: "Michigan",
				Zip: "12345",
			},
			hotel{
				Name: "West Hotel",
				Address: "456 Street",
				City: "California",
				Zip: "98765",
			},
		}
	)
}

}


