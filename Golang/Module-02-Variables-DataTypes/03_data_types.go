// Run: go run 03_data_types.go
package main

import "fmt"

func main() {
	// Integer types: int, int8, int16, int32, int64
	// Unsigned:      uint, uint8, uint16, uint32, uint64
	var i int = -42
	var u uint = 42
	fmt.Printf("int=%d uint=%d\n", i, u)

	// Float types: float32, float64 (default)
	var f float64 = 3.1415
	fmt.Printf("float=%f type=%T\n", f, f)

	// Boolean
	var done bool = true
	fmt.Println("done:", done)

	// String — immutable sequence of bytes (UTF-8).
	var greeting string = "नमस्ते Go"
	fmt.Println(greeting, "length(bytes):", len(greeting))

	// byte = alias for uint8 (one ASCII char)
	var b byte = 'A'
	fmt.Printf("byte=%d char=%c\n", b, b)

	// rune = alias for int32 (one Unicode code point)
	var r rune = '♥'
	fmt.Printf("rune=%d char=%c\n", r, r)

	// Complex numbers
	var c complex128 = complex(2, 3)
	fmt.Println("complex:", c)
}
