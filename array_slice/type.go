package array_slice

import (
	"fmt"
	"reflect"
)

func Run() {
	var arr = [3]int{1, 2, 3}
	arrtype := reflect.TypeOf(arr)

	fmt.Println("Type:", arrtype)        // Type: [3]int
	fmt.Println("Kind:", arrtype.Kind()) // Kind: array

	var sl = []int{1, 2, 3}
	sltype := reflect.TypeOf(sl)

	fmt.Println("Type:", sltype)        // Type: []int
	fmt.Println("Kind:", sltype.Kind()) // Kind: slice
}
