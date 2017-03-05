// Package main is the entry point of your application and must have
// the main() function.
package main

// Notice the import path: file path relative to $GOPATH
import "github.com/ebsarr/gintro/hello"

// All objects with names starting by a lowercase, like this main function,
// are unexported. They are only available to the package they reside in.
// Usually, unexported functions like this don't require documentation, as they
// will not be shown in the package documentation generated by godoc.
// You can still add internal comments accordingly.
// The main package must implement a main function.
func main() {
	// You don't need to specify the full path of the package to call it's functions.
	// However, if you imported two packages of the same name, you can call the fill
	// package name by including the package path.
	hello.SayHello()
}
