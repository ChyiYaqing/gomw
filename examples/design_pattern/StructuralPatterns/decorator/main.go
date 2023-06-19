package main

import (
	"fmt"
	"net/http"
)

/*
the Decorator Pattern
	> The Decorator Pattern is a design pattern that allows behavior to be added to an individual object,
	> dynamically, without affecting the behavior of other objects from the same class.
*/

type Printer interface {
	Print() string
}

type SimplePrinter struct{}

func (sp *SimplePrinter) Print() string {
	return "Hello, world!"
}

func BoldDecorator(p Printer) Printer {
	return PrinterFunc(func() string {
		return "<b>" + p.Print() + "</b>"
	})
}

type PrinterFunc func() string

func (pf PrinterFunc) Print() string {
	return pf()
}

/*
	middleware in web frameworks.
		> Middleware functions are functions that execute before or after a request is handled by a web server.
		> They can be used for tasks such as authentication, logging, and caching.
*/

type Handler func(r *http.Request) (int, string)

type Middleware func(h Handler) Handler

func LoggingMiddleware(h Handler) Handler {
	return func(r *http.Request) (int, string) {
		fmt.Println("Handling request...")
		status, body := h(r)
		fmt.Printf("Request handled with status %d\\\\n", status)
		return status, body
	}
}

func main() {
	simplePrinter := &SimplePrinter{}
	boldPrinter := BoldDecorator(simplePrinter)
	fmt.Println(simplePrinter.Print()) // Output: Hello, world!
	fmt.Println(boldPrinter.Print())   // Output: <b>Hello, world!</b>

	helloHandler := func(r *http.Request) (int, string) {
		return http.StatusOK, "Hello, world!"
	}
	loggingHandler := LoggingMiddleware(helloHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status, body := loggingHandler(r)
		w.WriteHeader(status)
		w.Write([]byte(body))
	})
	http.ListenAndServe(":1207", nil)
}
