package main

import (
	"os"
	"text/template"
	"log"
)

type car struct {
	Make, Model, Year string
}

type drivetrain struct {
	Cylinder string
	Cars []car
}

type engine struct {
	FWD, RWD, AWD drivetrain
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml.html"))
}

func main() {
	e := engine{
		FWD: drivetrain{
			Cylinder: "4-Cylinder",
			Cars: []car{
				car{"Honda", "Civic", "2016"},
				car{"Acura", "Integra", "2001"},
			},
		},
		AWD: drivetrain{
			Cylinder: "4-Cylinder",
			Cars: []car{
				car{"Mitsubishi", "Lancer Evolution", "2003"},
				car{"Subaru", "WRX STI", "2015"},
			},

		},
		RWD: drivetrain{
			Cylinder: "6-Cylinder",
			Cars: []car{
				car{"Toyota", "Supra", "1997"},
				car{"Nissan", "GTR", "2016"},
			},

		},
	}

	err := tpl.Execute(os.Stdout, e)
	if err != nil {
		log.Fatalln(err)
	}

}

/*
package main

import (
	"os"
	"text/template"
	"log"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "4"},
				course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				course{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming with Go", "5"},
				course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}

 */