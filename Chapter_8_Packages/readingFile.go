package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	// Reading a file using ioutil
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	// Make string from bytes to print
	str := string(bs)
	fmt.Println(str)
}
