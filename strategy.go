package main

import "fmt"

// PaymentStrategy defines the interface for payment methods.
type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card.\n", amount)
}


type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal.\n", amount)
}


type CashPayment struct{}

func (c *CashPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f in cash.\n", amount)
}


type PaymentContext struct {
	paymentStrategy PaymentStrategy
}

func (pc *PaymentContext) SetPaymentStrategy(strategy PaymentStrategy) {
	pc.paymentStrategy = strategy
}

func (pc *PaymentContext) ProcessPayment(amount float64) {
	pc.paymentStrategy.Pay(amount)
}

func main() {
	// Creating a payment context.
	paymentContext := &PaymentContext{}

	// Performing payments using different strategies.
	fmt.Println("Processing payments:")
	
	paymentContext.SetPaymentStrategy(&CreditCardPayment{})
	paymentContext.ProcessPayment(100.00)

	paymentContext.SetPaymentStrategy(&PayPalPayment{})
	paymentContext.ProcessPayment(50.25)

	paymentContext.SetPaymentStrategy(&CashPayment{})
	paymentContext.ProcessPayment(30.50)
}
