package main

import "fmt"

func main(){
	i := 1
	// Similar to other languages while loop
	for i <= 10 {
		fmt.Println(i)
		i = i+1
	}

	// For loop more traditional
	for x := 0; x <= 10; x++ {
		fmt.Println(x)
	}

	// If statement
	if i % 2 == 0 {
		// Even
	} else {
		// Odd
	}

	// Else if
	if i % 2 == 0 {
 		// divisible by 2
	} else if i % 3 == 0 {
 		// divisible by 3
	} else if i % 4 == 0 {
 		// divisible by 4
	}

	// Switch statement
	n := 3
	switch n {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
	}


}
