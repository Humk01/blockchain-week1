package main

import (
	"fmt"
	"crypto/sha256"
)

func main () {
	// simple hash

	data1 := "hello hashing"
	hash1 := sha256.Sum256([]byte(data1))

	fmt.Printf("Data: %s\n", data1)
    fmt.Printf("Hash: %x\n\n", hash1)

	data2 := "hello hashing"
	hash2 := sha256.Sum256([]byte(data2))

	fmt.Printf("Data: %s\n", data2)
    fmt.Printf("Hash: %x\n\n", hash2)
}