package main

import "fmt"

func main(){
	// Create array of int with size 5
	var x[5]int
	x[4] = 100
	fmt.Println(x)

  	// Create array of float64
	xs := [5]float64{
 		98,
 		93,
 		77,
 		82,
 		83,
	}

	// Add all values in array
	var total float64 = 0
	// Use _ when the index is not needed
	for _, value := range xs {
		total += value
	}
	fmt.Println(total / float64(len(xs)))

	// Creating a slice
	// Size of 5 with a capacity of 15
	s := make([]float64, 5, 15)
	s[2] = 10
	fmt.Println(s)

	// Append to a slice
	// Adding 4,5 to the end of slice s
	s = append(s, 4,5)
	fmt.Println(s)

	// Copy a slice to another
	s2 := make([]float64, 3)
	// Copy s into s2
	// Because s2 is smaller only the first 3 entries are copied
	copy(s2, s)
	fmt.Println(s2)

	// Creating a map
	// Use strings to map to ints
	mp := make(map[string]int)
	mp["key1"] = 101
	fmt.Println(mp["key1"])

	// Map of elements
	elements := map[string]map[string]string{
		 "H": map[string]string{
		 	"name":"Hydrogen",
		 	"state":"gas",
		 },
		 "He": map[string]string{
		 	"name":"Helium",
		 	"state":"gas",
		 },
		 "Li": map[string]string{
		 	"name":"Lithium",
		 	"state":"solid",
		 },
		 "Be": map[string]string{
		 	"name":"Beryllium",
		 	"state":"solid",
		 },
		 "B": map[string]string{
		 	"name":"Boron",
		 	"state":"solid",
		 },
		 "C": map[string]string{
		 	"name":"Carbon",
		 	"state":"solid",
		 },
		 "N": map[string]string{
			 "name":"Nitrogen",
		 	"state":"gas",
		 },
		 "O": map[string]string{
		 	"name":"Oxygen",
		 	"state":"gas",
		 },
		 "F": map[string]string{
		 	"name":"Fluorine",
		 	"state":"gas",
		 },
		 "Ne": map[string]string{
		 	"name":"Neon",
		 	"state":"gas",
		 },
 	}

	// Lookup returns two values the element (el) and if it was found (ok)
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}
