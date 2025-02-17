package designpattern

import "fmt"

// PaymentStrategy defines an interface for different payment methods.
type PaymentStrategy interface {
	Pay(amount float64)
}

// CreditCardStrategy is a concrete strategy for credit card payments.
type CreditCardStrategy struct {
	Name   string
	Number string
}

// Pay implements the PaymentStrategy interface for CreditCardStrategy.
func (c *CreditCardStrategy) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card (%s)\n", amount, c.Number)
}

// PaypalStrategy is a concrete strategy for PayPal payments.
type PaypalStrategy struct {
	Email string
}

// Pay implements the PaymentStrategy interface for PaypalStrategy.
func (p *PaypalStrategy) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal (%s)\n", amount, p.Email)
}

// ShoppingCart uses a PaymentStrategy to process payments.
type ShoppingCart struct {
	Items           []float64
	PaymentStrategy PaymentStrategy
}

// AddItem adds an item price to the shopping cart.
func (cart *ShoppingCart) AddItem(price float64) {
	cart.Items = append(cart.Items, price)
}

// Checkout calculates the total and uses the selected payment strategy.
func (cart *ShoppingCart) Checkout() {
	var total float64
	for _, price := range cart.Items {
		total += price
	}
	cart.PaymentStrategy.Pay(total)
}

func Run() {
	// Create a shopping cart and add items.
	cart := ShoppingCart{}
	cart.AddItem(29.99)
	cart.AddItem(49.50)

	// Use CreditCardStrategy to pay.
	cart.PaymentStrategy = &CreditCardStrategy{
		Name:   "John Doe",
		Number: "1234-5678-9012-3456",
	}
	cart.Checkout()

	// Switch strategy to PaypalStrategy.
	cart = ShoppingCart{} // Reset the cart for demonstration.
	cart.AddItem(99.99)
	cart.PaymentStrategy = &PaypalStrategy{
		Email: "john@example.com",
	}
	cart.Checkout()
}

// Explanation
// Interface Definition:
// The PaymentStrategy interface declares a Pay method. This allows us to define various payment methods that implement this interface.

// Concrete Strategies:

// CreditCardStrategy implements the payment method using credit card details.
// PaypalStrategy implements the payment method using an email address.
// Context (ShoppingCart):
// The ShoppingCart struct maintains a list of items and a field of type PaymentStrategy. Its Checkout method calculates the total and delegates the payment operation to the chosen strategy.

// Usage in Main:
// In main(), we create a shopping cart, add items, and set a specific payment strategy before calling Checkout(). We demonstrate switching strategies by resetting the cart and using a different payment method.

// This example demonstrates how the Strategy pattern allows the payment behavior to be changed dynamically at runtime, adhering to the Open/Closed Principle by keeping the shopping cart code independent of the payment algorithm implementations.
