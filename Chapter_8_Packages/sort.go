package main

import (
	"fmt"
	"sort"
)

// There are several sort functions for slices of ints and floats
// This is an example for custom data

type Person struct {
	Name string
	Age int
}

type ByName []Person

// Get the length
func (ps ByName) Len() int {
	return len(ps)
}

// Compare two values
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

// Swap two items places
func (ps ByName) Swap (i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main(){
	kids := []Person{
		{"Jill",9},
		{"Jack",10},
		{"Albert", 14},
		{"Zoey", 6},
	}

	// sort.Interface requires 3 methods Len, Less, Swap
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
