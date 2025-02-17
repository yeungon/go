package main

import (
	"fmt"
	"os"
	"reflect"
)

func ReadTheFile() {
	// Read entire file content into memory
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var b1 []byte  // Same as []uint8
	var b2 []uint8 // Same as []byte

	fmt.Println(reflect.TypeOf(b1)) // []uint8
	fmt.Println(reflect.TypeOf(b2)) // []uint8

	// Convert byte slice to string and print
	fmt.Println("File content:")

	fmt.Println(content)
	fmt.Println(reflect.TypeOf(content)) //[]uint8 --> slice of bytes
	fmt.Println(string(content))
}

func main() {
	fmt.Println("hello")
	ReadTheFile()
}
