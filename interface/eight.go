package interface_go

import "fmt"

// UserRepository is an interface that defines the methods for fetching user data.
type UserRepository8 interface {
	FindByID8(id int) string
}

// InMemoryUserRepository is a concrete implementation of UserRepository.
type InMemoryUserRepository8 struct{}

// FindByID returns a user name based on the ID (simulated in-memory data).
func (r *InMemoryUserRepository8) FindByID8(id int) string {
	return "User from Memory: John Doe"
}

// DatabaseUserRepository is another concrete implementation of UserRepository.
type DatabaseUserRepository8 struct{}

// FindByID returns a user name based on the ID (simulated database data).
func (r *DatabaseUserRepository8) FindByID8(id int) string {
	return "User from Database: Jane Smith"
}

// UserService depends on UserRepository to fetch user data.
type UserService8 struct {
	repo UserRepository8
}

// NewUserService is a constructor that injects the UserRepository dependency.
func NewUserService8(repo UserRepository8) *UserService8 {
	return &UserService8{repo: repo}
}

// GetUserName returns the user name for a given ID.
func (s *UserService8) GetUserName8(id int) string {
	return s.repo.FindByID8(id)
}

func Run8() {
	// Create instances of the different repository implementations.
	inMemoryRepo := &InMemoryUserRepository8{}
	databaseRepo := &DatabaseUserRepository8{}

	// Inject the InMemoryUserRepository into the UserService.
	userServiceWithMemory := NewUserService8(inMemoryRepo)
	fmt.Println(userServiceWithMemory.GetUserName8(1)) // Output: User from Memory: John Doe

	// Inject the DatabaseUserRepository into the UserService.
	userServiceWithDatabase := NewUserService8(databaseRepo)
	fmt.Println(userServiceWithDatabase.GetUserName8(1)) // Output: User from Database: Jane Smith
}
