package main

import ("fmt"
		"container/list")

func main() {
	// Create a new list called x
	var x list.List

	// Add values to list x using PushBack
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	// Get the Front item of the list then each next item until none are left
	// Lists are all doubly linked and have a Previous
	for e := x.Front(); e != nil; e=e.Next(){
		// Print the value of the item
		fmt.Println(e.Value.(int))
	}
}
