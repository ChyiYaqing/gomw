package main

/*
The Adaper Pattern
	1. Target Interface: the interface that the client is expecting to interact with.
	2. Adapter: the object that implements the Target Interface and wraps the adatee.
	3. Adaptee: the interface that needs to be adapted to be used by the client.
	4. Client: the component that uses the Target Interface.
*/

import "fmt"

type PayPal struct{}

func (p *PayPal) MakePayment(amount float32) bool {
	// connect to PayPal and process payment
	return true
}

type Amazon struct{}

func (a *Amazon) PayAmazon(amount float32) bool {
	// connect to Amazon and process payment
	return true
}

// interface that needs to be implemented by all the payment provides.
type PaymentGateway interface {
	ProcessPayment(amount float32) bool
}

type PayPalAdapter struct {
	PayPal *PayPal
}

func (p *PayPalAdapter) ProcessPayment(amount float32) bool {
	return p.PayPal.MakePayment(amount)
}

type AmazonAdapter struct {
	Amazon *Amazon
}

func (a *AmazonAdapter) ProcessPayment(amount float32) bool {
	return a.Amazon.PayAmazon(amount)
}

type Target interface {
	Request() string
}

// 适配者
type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct{}

func (a *adapteeImpl) SpecificRequest() string {
	return "Adaptee request"
}

type adapter struct {
	adaptee Adaptee
}

func (a *adapter) Request() string {
	return "Adapter: " + a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &adapteeImpl{}
	target := &adapter{
		adaptee: adaptee,
	}
	fmt.Println(target.Request())

	paymentGateway := &PayPalAdapter{PayPal: &PayPal{}}
	paymentGateway2 := &AmazonAdapter{Amazon: &Amazon{}}
	// triggers Paypal
	paymentGateway.ProcessPayment(100)
	// triggers Amazon
	paymentGateway2.ProcessPayment(100)

}
