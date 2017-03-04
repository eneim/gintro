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
	"fmt"
	"os"
)

// In this example, our package is the main package and must have a main() function.
// The main() function takes no argument and returns nothing.
func main() {
	var name string
	if len(os.Args) >= 2 {
		name = os.Args[1]
	} else {
		name = "World"
	}
	fmt.Printf("Hello, %s!\n", name)
}
