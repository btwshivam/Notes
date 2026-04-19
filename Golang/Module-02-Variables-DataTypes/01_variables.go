// Run: go run 01_variables.go
package main

import "fmt"

// Package-level variable (accessible everywhere in this package).
var version = "1.0.0"

func main() {
	// 1. Explicit type + value
	var name string = "Go"
	fmt.Println("Name:", name)

	// 2. Type inferred from value
	var age = 15
	fmt.Println("Age:", age)

	// 3. Short declaration (only inside functions)
	year := 2009
	fmt.Println("Year:", year)

	// Multiple variables in one line
	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

	// Grouped declaration
	var (
		isActive bool   = true
		count    int    = 5
		label    string = "Gophers"
	)
	fmt.Println(isActive, count, label)

	// Zero values: variables declared without value get default zero.
	var (
		zInt    int
		zFloat  float64
		zStr    string
		zBool   bool
	)
	fmt.Printf("zero values -> int:%d float:%f str:%q bool:%t\n",
		zInt, zFloat, zStr, zBool)

	// Variables can be reassigned, but NOT re-declared in the same scope.
	age = 16
	fmt.Println("Updated age:", age)
	fmt.Println("Version:", version)
}
