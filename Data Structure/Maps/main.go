package main

import "fmt"

func main() {
	/*how to create a map
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
	*/

	/*
	myGreeting := map[int]string{
		0: "Good Morning",
		1: "good Afternoon",
		2: "Good Evening",
	}

	fmt.Println(myGreeting, 1)
*/
	//composite literal
	words := map[string]string{
		"hush" : "be quiet",
		"0" :	"a number; a loser; emptiness; enlightenment; off",
		"golang" : "the best programming language ever",
	}

	//finish this program by using range to iterate over the map, printing out the key and value
	for k, v := range words {
		fmt.Println(k, ":", v)
	}

	//delete an entry
	delete(words, "hush")


	//add a value
	words["cowboy"] = "rides horses"
}