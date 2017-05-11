package main

import "fmt"

func main(){
	// Use defer to ensure recover runs
	defer func(){
		// Recover stops the panic and returns the value passed to panic
		str := recover()
		fmt.Println(str,"recovered")
	}()
	panic("Panic")
}
