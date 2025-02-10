package interface_go

import (
	"fmt"
)

// Define the Animal interface
type Animal interface {
	Speak() string
	Move() string
}

// Define the Pet interface
type Pet interface {
	Owner() string
}

// Dog struct
type Dog struct {
	owner string
}

// Cat struct
type Cat struct {
	owner string
}

// Implement Speak method for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Implement Move method for Dog
func (d Dog) Move() string {
	return "Runs on four legs"
}

// Implement Owner method for Dog
func (d Dog) Owner() string {
	return d.owner
}

// Implement Speak method for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Implement Move method for Cat
func (c Cat) Move() string {
	return "Jumps gracefully"
}

// Implement Owner method for Cat
func (c Cat) Owner() string {
	return c.owner
}

// PetService struct depends on interfaces instead of concrete types
type PetService struct {
	animal Animal
	pet    Pet
}

// NewPetService is a constructor that takes dependencies as interfaces
func NewPetService(animal Animal, pet Pet) *PetService {
	return &PetService{animal: animal, pet: pet}
}

// Describe function uses dependency injection to describe the pet
func (ps *PetService) Describe() {
	fmt.Println("This animal says:", ps.animal.Speak())
	fmt.Println("It moves like this:", ps.animal.Move())
	fmt.Println("This pet belongs to:", ps.pet.Owner())
}

func Run4() {
	// Inject dependencies (Dog and Cat) into PetService
	dog := Dog{owner: "Alice"}
	cat := Cat{owner: "Bob"}

	// Using dependency injection
	dogService := NewPetService(dog, dog)
	catService := NewPetService(cat, cat)

	fmt.Println("Dog Service:")
	dogService.Describe()

	fmt.Println("\nCat Service:")
	catService.Describe()
}

// How This Demonstrates Dependency Injection:
// Interfaces (Animal and Pet) as Dependencies

// The PetService struct does not depend on concrete implementations (Dog or Cat).
// It only depends on the Animal and Pet interfaces, making it flexible and reusable.
// Constructor Injection (NewPetService)

// Dependencies are injected into PetService when calling NewPetService(animal Animal, pet Pet).
// This allows for easy swapping of implementations.
// Decoupling Business Logic from Implementation

// PetService does not know whether it's working with Dog or Cat, making it highly testable and flexible.
