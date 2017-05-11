package main

import "fmt"

func main(){
	// Declare and assign variables
	var x = "Hello World"
	fmt.Println(x)

	// Compare variables
	var z string = "hello"
	var y string = "world"
	fmt.Println(z == y) // false

	// Declare and assign variables 2
	a := "Hello, world"
	fmt.Println(a)

	// Defining names
	name := "Max"
	fmt.Println(name)

	dogsName := "Max"	// Camel case
	fmt.Println(dogsName)

	// Define multiple variables
	var (
		q = 5
		r = 10
		s = 15
	)
	fmt.Println(q,r,s)

}
