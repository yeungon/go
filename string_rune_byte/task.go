package lt_string

import "fmt"

// extract the character

func Task() {
	s := "Chào bạn"
	fmt.Println(s[2]) //Wrong way

	runes := []rune(s)
	fmt.Println(string(runes[2]))
}
