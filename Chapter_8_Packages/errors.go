package main

import ("errors"
		"fmt")

func main(){
	// Create a new error
	err := errors.New("Error Message")
	// Print it
	fmt.Println(err)
}
