package interface_go

import "fmt"

// UserRepository is an interface that defines the methods for fetching user data.
type UserRepository interface {
	FindByID(id int) string
}

// InMemoryUserRepository is a concrete implementation of UserRepository.
type InMemoryUserRepository struct{}

// FindByID returns a user name based on the ID.
func (r *InMemoryUserRepository) FindByID(id int) string {
	// In a real application, this would fetch data from a database or another source.
	return "John Doe"
}

// UserService depends on UserRepository to fetch user data.
type UserService struct {
	repo UserRepository
}

// NewUserService is a constructor that injects the UserRepository dependency.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUserName returns the user name for a given ID.
func (s *UserService) GetUserName(id int) string {
	return s.repo.FindByID(id)
}

func Run7() {
	// Create an instance of the InMemoryUserRepository.
	repo := &InMemoryUserRepository{}

	// Inject the repository into the UserService.
	userService := NewUserService(repo)

	// Use the UserService to get a user name.
	userName := userService.GetUserName(1)
	fmt.Println("User Name:", userName) // Output: User Name: John Doe
}

// This code demonstrates dependency injection in Go, where the UserService depends on a UserRepository interface,
//and the concrete implementation (InMemoryUserRepository) is injected at runtime.
//Explanation
// UserRepository Interface: We define an interface UserRepository that specifies the FindByID method. This allows us to decouple the UserService from the concrete implementation of the repository.

// InMemoryUserRepository: This is a concrete implementation of the UserRepository interface. It simulates fetching user data from an in-memory source.

// UserService: The UserService depends on the UserRepository interface. The dependency is injected via the constructor NewUserService.

// main Function: In the main function, we create an instance of InMemoryUserRepository and inject it into the UserService. This allows us to easily swap out the repository implementation without changing the UserService.

// Benefits of Dependency Injection
// Testability: You can easily mock the UserRepository in unit tests.

// Flexibility: You can switch between different implementations of UserRepository without modifying the UserService.

// Decoupling: The UserService doesn't need to know about the concrete implementation of UserRepository.

// This is a simple example, but the same principles apply to more complex applications.
