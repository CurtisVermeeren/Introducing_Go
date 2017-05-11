package main

import (
	"fmt"
	"hash/crc32"
)


func main() {
	// Create a hasher
	h := crc32.NewIEEE()
	// Wrtie data to it
	h.Write([]byte("test"))
	// Calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}
