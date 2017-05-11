package main

import "fmt"

// Pointer represented by * followed by a type
// * is used to dereference. This allows access to the value of the pointer
func zero(xPtr *int){
	*xPtr = 0
}

func main(){
	x := 5
	// Use & to find the address of a variable
	// &x returns a *int because x is an int
	zero(&x)
	fmt.Println(x)
	fmt.Println("The memory address of x is",&x)
}
