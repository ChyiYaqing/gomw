package main

import "fmt"

/*
the Command Pattern
*/

// 1. Command: an interface that defines the execution method.
type Command interface {
	Execute() string
}

// 2. ConcreteCommand: a struct that implements the Command interface and contains a receiver object.
type ConcreteCommand struct {
	receiver Receiver
}

func (cc *ConcreteCommand) Execute() string {
	return cc.receiver.Action()
}

// 3. Invoker: a struct that triggers the execution of the command.
type Invoker struct {
	command Command
}

func (i *Invoker) ExecuteCommand() string {
	return i.command.Execute()
}

// 4. Receiver: a struct that receives the request and performs the action.
type Receiver struct{}

func (r *Receiver) Action() string {
	return "Action Performed"
}

func main() {
	receiver := &Receiver{}
	concreteCommand := &ConcreteCommand{receiver: *receiver}
	invoker := &Invoker{command: concreteCommand}
	result := invoker.ExecuteCommand()
	fmt.Println(result)
}
