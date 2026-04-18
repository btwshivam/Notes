// Run: go run 03_comments.go

// Single-line comment — use for short notes.

/*
   Multi-line comment.
   Useful for longer explanations or temporarily disabling blocks of code.
*/

// Package-level doc comment appears in `go doc` output.
package main

import "fmt"

// Greet prints a greeting. A comment that starts with the identifier
// name becomes its official documentation.
func Greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	Greet("Students")
	// TODO: Add more greetings later.
}
