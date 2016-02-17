package main
import "fmt"

func main() {

	//make slice of int, using composite literal, and print it.
	//Afterwards,  print the index and value using a range statement.

	x := []int{2,4,6,8}
	fmt.Println(x)
	//length of the slice
	fmt.Println(len(x))
	fmt.Println("-------------------------")

	for i, v := range x {
		fmt.Println(i, v)
	}


	//slice example

	/*
	var x []int
	x = append(x, 5)
	fmt.Println(x)
	*/

	/*
	x:= make([]int, 0, 4)
	x[0] = 5
	x[1] =
	x = append(x, 5)
	a(x)
	fmt.Println(x)
	*/

	/*
	var y int
	fmt.Scanln(&y)
	x := []int{4,6,7,9,5}
	x = x[y:4]

	fmt.Println(x)
	*/

	//slice is a pointer to an array
	//header, length, capacity
	//composite literal
	//y := make([]int, 5, 9)

	/*
	x := []int{4,6,7,9,5}
	x = append(x, 6, 8, 0, 2)
	smallest := x[0]
	for _, v := range x {
		if v < smallest {
			smallest = v
		}
	}
	*/

}

/*
func a(t []int) {
	panic("Error")
*/

}