package main

import "fmt"

/*
the Mediator Pattern : 调解员模式
	> To implement the Mediator Pattern in Golang, we wil start by defining an
	interface for our mediator object. This interface will define the methods that
	our mediator object will use to communicate with the different objects in our system.

the Mediator Pattern vs the Facade Pattern
	1. the Mediator Pattern focuses on simplifying communication between objects
	2. the Facade Pattern focuses on simplifying the interface to a larger body of code.
*/

// 调解员
type mediator interface {
	addUser(user user)
	sendMessage(sender user, message string)
}

type user struct {
	id   int
	name string
}

type chatRoomMediator struct {
	users []user
}

func (c *chatRoomMediator) addUser(user user) {
	c.users = append(c.users, user)
}

func (c *chatRoomMediator) sendMessage(sender user, message string) {
	for _, u := range c.users {
		if u.id != sender.id {
			fmt.Printf("%s sent message to %s: %s\n", sender.name, u.name, message)
		}
	}
}

func main() {
	mediator := &chatRoomMediator{}
	user1 := user{id: 1, name: "user1"}
	user2 := user{id: 2, name: "user2"}
	user3 := user{id: 3, name: "user3"}
	mediator.addUser(user1)
	mediator.addUser(user2)
	mediator.addUser(user3)
	mediator.sendMessage(user1, "Hello, everyone!")
}
