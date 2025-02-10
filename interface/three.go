package interface_go

// import (
// 	"fmt"
// )

// // Define the Animal interface with two methods
// type Animal interface {
// 	Speak() string
// 	Move() string
// }

// // Define the Pet interface with an additional method
// type Pet interface {
// 	Owner() string
// }

// // Dog struct
// type Dog struct {
// 	owner string
// }

// // Cat struct
// type Cat struct {
// 	owner string
// }

// // Implement Speak method for Dog
// func (d Dog) Speak() string {
// 	return "Woof!"
// }

// // Implement Move method for Dog
// func (d Dog) Move() string {
// 	return "Runs on four legs"
// }

// // Implement Owner method for Dog
// func (d Dog) Owner() string {
// 	return d.owner
// }

// // Implement Speak method for Cat
// func (c Cat) Speak() string {
// 	return "Meow!"
// }

// // Implement Move method for Cat
// func (c Cat) Move() string {
// 	return "Jumps gracefully"
// }

// // Implement Owner method for Cat
// func (c Cat) Owner() string {
// 	return c.owner
// }

// // DescribeAnimal takes an Animal interface and prints its behaviors
// func DescribeAnimal(a Animal) {
// 	fmt.Println("This animal says:", a.Speak())
// 	fmt.Println("It moves like this:", a.Move())
// }

// // DescribePet takes a Pet interface and prints owner information
// func DescribePet(p Pet) {
// 	fmt.Println("This pet belongs to:", p.Owner())
// }

// func Run3() {
// 	// Create instances of Dog and Cat
// 	dog := Dog{owner: "Alice"}
// 	cat := Cat{owner: "Bob"}

// 	// Call DescribeAnimal function for both Dog and Cat
// 	fmt.Println("Dog:")
// 	DescribeAnimal(dog)
// 	DescribePet(dog)

// 	fmt.Println("\nCat:")
// 	DescribeAnimal(cat)
// 	DescribePet(cat)
// }

// We introduced a new interface called Pet, which requires an Owner() method.
// Both Dog and Cat now have an owner field and implement Owner().
// The function DescribePet(p Pet) is added to display the owner's name.
// In main(), we create Dog and Cat instances with owners and pass them to DescribeAnimal() and DescribePet().
