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

func ExamplePage() {
	checkIns := map[string]bool {
		"Alex": true,
		"Brett": false,
		"Tustin": true,
		"Ziggy": false,
		"Emma": false,
	}

	Page(checkIns)

	// Unordered output:
	// Paging Brett; please see the front desk to check in.
	// Paging Ziggy; please see the front desk to check in.
	// Paging Emma; please see the front desk to check in.
}