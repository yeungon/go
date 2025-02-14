package interface_go

import "fmt"

func Run() {
	var notype interface{}
	notype = "xin chào, hello world"
	fmt.Println(notype)
	notype = 123.5 /// int

	fmt.Println(notype)

}
