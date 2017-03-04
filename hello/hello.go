// Every file belongs to a package. The package `main` is the entry point of
// your application and must contain the main() function.
// Usually, this comment section should have a bried explaination of what your
// package does, and some additional information as needed. You can consult
// the golang.org page for more information on coding style.
package main

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

// In this example, our package is the main package and must have a main() function.
// The main() function takes no argument and returns nothing.
func main() {
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
