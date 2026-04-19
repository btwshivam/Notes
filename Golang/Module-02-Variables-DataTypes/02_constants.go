// Run: go run 02_constants.go
package main

import "fmt"

// Single constant
const Pi = 3.14159

// Typed constant
const MaxUsers int = 100

// Grouped constants
const (
	AppName    = "MyApp"
	AppVersion = "2.1.0"
	MaxRetries = 3
)

func main() {
	fmt.Println("Pi:", Pi)
	fmt.Println("Max users:", MaxUsers)
	fmt.Println(AppName, AppVersion, "retries:", MaxRetries)

	// Constants must be known at compile time.
	// The line below would FAIL: const now = time.Now()

	// Constants are untyped until used — they take the type of the context.
	const n = 42
	var i int = n
	var f float64 = n
	fmt.Println(i, f)
}
