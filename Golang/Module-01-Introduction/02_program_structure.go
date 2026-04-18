// Run: go run 02_program_structure.go
//
// Every Go source file has THREE main sections:
//   1. package declaration  (required)
//   2. imports              (optional)
//   3. top-level code       (funcs, vars, types, consts)

package main

// Grouped import — preferred when importing multiple packages.
import (
	"fmt"
	"math"
)

// A top-level (package-level) variable.
var appName = "StructureDemo"

// main is the entry point of any executable Go program.
func main() {
	fmt.Println("App:", appName)
	fmt.Println("Square root of 144 =", math.Sqrt(144))
}
