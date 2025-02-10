package interface_go

// import "fmt"

// // Define the Animal interface with two methods
// type Animal interface {
// 	Speak() string
// 	Move() string
// }

// // Dog struct
// type Dog struct{}

// // Cat struct
// type Cat struct{}

// // Implement Speak method for Dog
// func (d Dog) Speak() string {
// 	return "Woof!"
// }

// // Implement Move method for Dog
// func (d Dog) Move() string {
// 	return "Runs on four legs"
// }

// // Implement Speak method for Cat
// func (c Cat) Speak() string {
// 	return "Meow!"
// }

// // Implement Move method for Cat
// func (c Cat) Move() string {
// 	return "Jumps gracefully"
// }

// // DescribeAnimal takes an Animal interface and prints its behaviors
// func DescribeAnimal(a Animal) {
// 	fmt.Println("This animal says:", a.Speak())
// 	fmt.Println("It moves like this:", a.Move())
// }

// func Run2() {
// 	// Create instances of Dog and Cat
// 	var dog Animal = Dog{}
// 	var cat Animal = Cat{}

// 	// Call DescribeAnimal function for both Dog and Cat
// 	fmt.Println("Dog:")
// 	DescribeAnimal(dog)

// 	fmt.Println("\nCat:")
// 	DescribeAnimal(cat)
// }

// // The Animal interface now has two methods: Speak() and Move().
// // Both Dog and Cat implement the Move() method in addition to Speak().
// // We define a function DescribeAnimal(a Animal) that takes an interface as a parameter and calls both methods on it.
// // The main() function demonstrates how we can use the Animal interface to interact with different types in a unified way.
