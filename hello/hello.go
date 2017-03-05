// Package hello provides a single function to implements `hello, world!` program in go.
//
// As a convention, package documentation starts with "Package packageName" followed by a brief
// description of what the package does in one line.
// You could document further by added more lines of comments before the package name.
package hello

// imports can be grouped like this.
// You can also use the import keyword without grouping:
//	import "packageA"
//	import "packageB"
import (
	// fmt implements formatted I/O with functions analogous to C's printf and scanf
	// Try to be familiar with this package as you'll definitely use it frequently
	"fmt"

	// os provides a platform-independent interface to operating system functionality
	"os"
)

// SayHello prints "hello, world!"
// As a convention, doc strings start with the name of the funtion  you are documenting.
func SayHello() {
	// Use the var keyword to declare a variable.
	// A newly declared varialbe is assigne the zero value of the type it represents.
	// In this case, name has the string's zero value "".
	// Inside a function, we can implicitly declare the type of a variable
	// with the := operator.
	//		name := ""
	// Since "" is the value of a string, go will implicitly declare name as type string
	// Another example: Suppose we have the following function:
	//		func getString() string {
	// 			return "I am a string!"
	//		}
	// You could use the := operator to implicitly declare a variable of type string
	// and assign the returned value of the function, provided you do it inside a function
	//		str := getString()
	var name string
	if len(os.Args) >= 2 {
		// os.Args is a slice of strings. you can do this assignment only if name was
		// previously declared as type string. go is a statically typed language!!!
		name = os.Args[1]
	} else {
		name = "World"
	}
	fmt.Printf("Hello, %s!\n", name)
}
