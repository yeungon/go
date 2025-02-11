package interface_go

import "fmt"

func Run() {

	var notype interface{}
	notype = "type string"
	fmt.Println(notype)
	notype = 123

	fmt.Println(notype)

}
