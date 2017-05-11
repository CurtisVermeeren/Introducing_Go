package main

import "fmt"

// Function to compute an average of an array of floats
// Takes xs an array of float64 and returns a single float64
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Function to add values
// The ... before the type means the function can take 0 or more of that type
func add(args ...int) int{
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// Using closure to return a function inside another function
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}

// A recursive factorial function calls itself
func factorial(x uint) uint {
	if (x == 0){
		return 1
	}
	return x * factorial(x-1)
}

func main(){
	xs := []float64{98,93,77,82,83}

	// Possible to create functions in functions
	add2 := func(x,y int) int {
		return x + y
	}

	fmt.Println(average(xs))
	fmt.Println(add(1,2,3))
	fmt.Println(add2(1,2))

	x := 0

	// A function using a non local variable is known as closure
 	increment := func() int {
 		x++
 		return x
 	}

 	fmt.Println(increment())
 	fmt.Println(increment())

	// Call the closure number generator
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4


	// Call the recursive factorial
	fmt.Println(factorial(4))
}
