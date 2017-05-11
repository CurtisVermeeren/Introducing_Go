package main

import "fmt"

func first(){
	fmt.Println("1st")
}

func second(){
fmt.Println("second")
}

func main(){
	// Defer will call the function after the current function completes no matter what
	defer second()
	first()
}
