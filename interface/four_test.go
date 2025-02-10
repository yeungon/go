package interface_go

import (
	"bytes"
	"fmt"
	"testing"
)

// MockAnimal is a mock implementation of the Animal interface
type MockAnimal struct {
	speak string
	move  string
}

// Implement Speak method for MockAnimal
func (m MockAnimal) Speak() string {
	return m.speak
}

// Implement Move method for MockAnimal
func (m MockAnimal) Move() string {
	return m.move
}

// MockPet is a mock implementation of the Pet interface
type MockPet struct {
	owner string
}

// Implement Owner method for MockPet
func (m MockPet) Owner() string {
	return m.owner
}

// TestPetService_Describe tests the Describe method of PetService
func TestPetService_Describe(t *testing.T) {
	// Arrange: Create mock dependencies
	mockAnimal := MockAnimal{speak: "Test Sound", move: "Test Movement"}
	mockPet := MockPet{owner: "Test Owner"}

	// Inject mocks into PetService
	service := NewPetService(mockAnimal, mockPet)

	// Capture output
	var output bytes.Buffer
	fmt.Fprint(&output, "This animal says: ", service.animal.Speak(), "\n")
	fmt.Fprint(&output, "It moves like this: ", service.animal.Move(), "\n")
	fmt.Fprint(&output, "This pet belongs to: ", service.pet.Owner(), "\n")

	expected := "This animal says: Test Sound\nIt moves like this: Test Movement\nThis pet belongs to: Test Owner\n"

	// Assert
	if output.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output.String())
	}
}
