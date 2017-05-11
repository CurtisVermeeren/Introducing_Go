package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	/*
    Os: Some useful functions from the os package
	*/

	// Open a file
	file, err := os.Open("test.txt")
	if err != nil {
		// Handle errors here
		return
	}

	// Close the file at the end of main
	defer file.Close()

	// Get file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// Reading the file into byte slice
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	// Convert bytes to string and print
	str := string(bs)
	fmt.Println(str)

	// Create a file
	file2, err2 := os.Create("test2.txt")
	if err2 != nil {
		return
	}
	// Close the file at the end of main
	defer file2.Close()

	// Writing to the file
	file2.WriteString("test!")

	// Getting the contents of a directory use open on a directory
	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

 	// Read infro from the directory
	// Pass a limit of entries to Readdir. -1 means read all
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	// Print all the directory info
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// The function passed will run for every file and folder in the root folder
	// In this case the root is .
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
