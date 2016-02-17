package main

import "fmt"

type person struct{
	name string
	age int
	height int
}

func main() {
	x := person{
		name: "David",
		age: 21,
	}
	x.height = 6
	fmt.Println(x)
}

/*
type Person struct{
	name string
	age int
	height int
}

func (p Person) String() string{
	return fmt.Sprintf("Name: %x, Age: %d, Height: %v", p.name, p.age, p.height)

}

*/

/*
//create a type that holds a person's name, and the pets they own.

type person struct{
	name string
	pets []string

}

func main() {
	x := person{
		name: "David",
		pets: []string{"dog1", "dog2"},
	}

	fmt.Println(x)
}
*/

/*
package main

import "fmt"


type candyStore struct{
	Name string
	Candies []string
	}
//struct literal
//composite literal
func main() {
	st1 := candyStore{
		Name: "Circle K",
		Candies: []string{"Snickers", "M&M's",},
	}
	fmt.Println(st1)
}

*/

/*
package main

import "fmt"


type candyStore struct{
	Name string
	Candies map[string]int
	}
//struct literal
//composite literal
func main() {
	st1 := candyStore{
		Name: "Circle K",
		Candies: map[string]int{
			"Snickers": 3,
			"M&M's": 7,
	},
	fmt.Println(st1)
	fmt.Println(st1.Candies["Snickers"])
}

 */