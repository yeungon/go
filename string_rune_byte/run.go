package lt_string

import (
	"fmt"
	"os"
	"reflect"
	"unicode/utf8"
)

func Length() {
	s := "Hello, xin chào 😄"
	fmt.Println(len(s))                    // 21
	fmt.Println(utf8.RuneCountInString(s)) // 17
	runes := []rune(s)
	fmt.Println(len(runes)) //17
}

func ReadTheFile() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	content, err := os.ReadFile(dir + "/string_rune_byte/example.txt")
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

func Run() {
	ReadTheFile()
	Length()
	Task()
	Index()
}
