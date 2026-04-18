// Run: go run 04_fmt_basics.go
//
// The `fmt` package handles formatted I/O — the workhorse of any Go program.

package main

import "fmt"

func main() {
	// Println — prints values separated by spaces, adds newline.
	fmt.Println("Hello", "Go", 2026)

	// Print — no newline, no separator.
	fmt.Print("A")
	fmt.Print("B\n")

	// Printf — formatted output using verbs.
	name := "Alice"
	age := 30
	pi := 3.14159
	fmt.Printf("Name: %s, Age: %d, Pi: %.2f\n", name, age, pi)

	// Sprintf — returns a formatted string instead of printing it.
	msg := fmt.Sprintf("User %s is %d years old", name, age)
	fmt.Println(msg)

	// Common verbs:
	//   %d  integer
	//   %f  float
	//   %s  string
	//   %t  bool
	//   %v  default format for any type
	//   %T  type of the value
	//   %q  quoted string
	fmt.Printf("%v is of type %T\n", pi, pi)

	// Reading from user — Scanln.
	var input string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&input)
	fmt.Println("Welcome,", input)
}
