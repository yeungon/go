package interface_go

// import "fmt"

// // Define the Animal interface with a Speak method
// type Animal interface {
// 	Speak() string
// }

// // Dog struct
// type Dog struct{}

// // Cat struct
// type Cat struct{}

// // Implement the Speak method for Dog
// func (d Dog) Speak() string {
// 	return "Woof!"
// }

// // Implement the Speak method for Cat
// func (c Cat) Speak() string {
// 	return "Meow!"
// }

// func Run() {
// 	// Create instances of Dog and Cat
// 	var dog Animal = Dog{}
// 	var cat Animal = Cat{}

// 	// Call the Speak method on each
// 	fmt.Println(dog.Speak()) // Output: Woof!
// 	fmt.Println(cat.Speak()) // Output: Meow!
// }

// In this example:

// We define an Animal interface with a single method Speak that returns a string.
// We create two structs, Dog and Cat.
// Both structs implement the Speak method, thus satisfying the Animal interface.
// In the main function, we create instances of Dog and Cat, assign them to variables of type Animal, and call their Speak methods.
// This demonstrates how interfaces in Go allow different types to be used interchangeably as long as they implement the required methods.
