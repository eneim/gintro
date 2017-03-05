// Package hello provides a single function to implements `hello, world!` program in go.
package hello

import (
	"fmt"
	"os"
)

// Person represents an individual.
// struct is the basic type you can define in Go.
// it's a type that can hold any number of attributes and you can implement
// methods on it.
type Person struct {
	// An attribute that start Uppuer case is exported.
	FirstName string
	LastName  string
}

// SetFirstName assigns a first name to Person
// function with receivers of types Person or *Person are methods
// of the type Person.
// Note that Person and *Person are two differents types, and offer
// different behaviours when used as receivers:
//
//	Person as a receiver means that the method receives a value of the type Person.
//  When the function is called, a copy of the value will be passed to the function,
//	As a result, any change you do on the value inside the method will not affect
// 	or will not be seen by the caller.
//  On the other hand
// 	*Person as a receiver means a pointer is passed to the function. Any change you
//	do inside the function will be seen by the caller.
func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

// SetLastName assigns a last name to Person
func (p *Person) SetLastName(lastName string) {
	p.FirstName = lastName
}

// GetFirstName returns the first name
func (p Person) GetFirstName() string {
	return p.FirstName
}

// GetLastName returns the last name
func (p Person) GetLastName() string {
	return p.LastName
}

func (p Person) String() string {
	return fmt.Sprintf("Person: %s %s", p.GetFirstName(), p.GetLastName())
}

// Doctor represents an individual who can heal Persons
type Doctor struct {
	*Person
}

func (d Doctor) String() string {
	return fmt.Sprintf("Doctor: %s %s", d.GetFirstName(), d.GetLastName())
}

// SayHello prints "hello, world!"
func SayHello() {
	var name string
	if len(os.Args) >= 2 {
		name = os.Args[1]
	} else {
		name = "World"
	}
	fmt.Printf("Hello, %s!\n", name)
}
