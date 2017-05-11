package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
    Strings: Some useful functions from the strings package
	*/

	// Check if a string contains another
	fmt.Println(strings.Contains("test", "es")) // true

	// Count number of times a string occurs in another
	fmt.Println(strings.Count("test", "t")) // 2

	// Check if string is a prefix for another
	fmt.Println(strings.HasPrefix("text", "te")) // true

	// Check if string is a suffic for another
	fmt.Println(strings.HasSuffix("test", "st")) // true

	// Get the index of a string in another. -1 if doesn't exist
	fmt.Println(strings.Index("text", "e")) // 1

	// Join a list of strings together around another string
	fmt.Println(strings.Join([]string{"a","b","c"}, "-")) // a-b-c

	// Repeat a string a number of times
	fmt.Println(strings.Repeat("Ho ",3)) // Ho Ho Ho

	// Split a string into a list based on a string
	fmt.Println(strings.Split("a-b-c-d-e", "-")) // []string{"a","b","c","d","e"}

	// Replace a part of a string with another string x amount of times
	fmt.Println(strings.Replace("aaaa","a","b",2)) // bbaa

	// Convert to all lowercase and all uppercase
	fmt.Println(strings.ToLower("LOWER")) // lower
	fmt.Println(strings.ToUpper("upper")) // UPPER

	// Convert a string to a slice of bites and back
	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t','e','s','t'})
	fmt.Println(str)
}
