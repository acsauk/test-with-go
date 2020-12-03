// Package example is a package for demonstrating examples in source code
package examples

import (
	"fmt"
)

type Demo struct {}

func (d Demo) Hello() {}

// Hello prints out hello to the person provided
func ExampleHello() {
	greeting := Hello("Alex")
	fmt.Println(greeting)

	// Output:
	// Hello, Alex
}
