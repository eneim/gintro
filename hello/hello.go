// Package hello provides a single function to implement `hello, world!` program in go
package hello

import (
	"fmt"
	"os"
)

// Age represents the age of a Person
//
// We can also create our own types based on current types
// Here we create the type Age as an int.
type Age int

// SetAge sets the age
// We can now add methods to our type just like we did with structs
func (a *Age) SetAge(age int) {
	*a = Age(age)
}

// GetAge returns the age
func (a Age) GetAge() int {
	return int(a)
}

func (a Age) String() string {
	return fmt.Sprintf("%d", int(a))
}

// Person represents an individual.
// struct is the basic type you can define in Go.
// it's a type that can hold any number of attributes and you can implement
// methods on it.
type Person struct {
	// An attribute that start Uppuercase is exported.
	FirstName string
	LastName  string
	MyAge     Age
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
	p.LastName = lastName
}

// GetFirstName returns the first name
// *Person is a pointer to Person, but we could still call this method on
// a pointer: That is:
//	var p1 Person
//  p1 = Person{FirstName: Bassirou, LastName: Sarr}
//  var p2 *Person
//	p2 = &Person{FirstName: Bassirou, LastName: Sarr}
//  firstname1 := p1.GetFirstName() // This works
//  firstname2 := p2.GetFirstName() // This works as well
// What happens under the hood is that Go will automatically convert the received
// pointer into a value.
// In this case, the pointer is automatically be processed like return (&p).FirstName
func (p Person) GetFirstName() string {
	return p.FirstName
}

// GetLastName returns the last name
func (p Person) GetLastName() string {
	return p.LastName
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s", p.GetFirstName(), p.GetLastName())
}

// Doctor represents an individual who can heal Persons
type Doctor struct {
	// Go does not support sub-classing, but you could embed a type inside
	// another type, just like we do here:
	// We embed type *Person in type Doctor -> Doctor has all the attributes
	// and methods that *Person has
	// We can now add other attributes to Doctor, and implement other methods
	// This is not quite like sub-classing, but that's all you get in go
	*Person
}

// Doctor does have anything to do with Person, although we embed Person in Doctor
// Doctor and Person both have the same method String(), but it's not overriding.
func (d Doctor) String() string {
	return fmt.Sprintf("Dr. %s", d.GetLastName())
}

// SayHello ...
// Let's try it out....
func SayHello() {
	var firstName string
	var lastName string
	if len(os.Args) == 3 {
		firstName = os.Args[1]
		lastName = os.Args[2]
	} else {
		firstName = "Bassirou"
		lastName = "Sarr"
	}

	// let's create a person
	// In fact, we don't even need to create setters or getters for our objects
	// We made them accessible by naming them with starting Uppuercase letters.
	//
	p := &Person{
		FirstName: firstName,
		LastName:  lastName,
	}
	p.MyAge.SetAge(35)

	d := &Doctor{
		&Person{},
	}
	d.SetFirstName(firstName)
	d.SetLastName(lastName)
	d.MyAge.SetAge(35)

	fmt.Printf("Hello, %s. Your age is %s.\n", p, p.MyAge)
	fmt.Printf("Hello, %s. Your age is %s.\n", d, d.MyAge)
}
