// Run: go run 04_type_conversion.go
//
// Go has NO implicit type conversion — you must convert explicitly.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int → float
	i := 10
	f := float64(i)
	fmt.Printf("%d -> %f\n", i, f)

	// float → int (truncates the decimal, does not round)
	x := 3.99
	y := int(x)
	fmt.Printf("%f -> %d\n", x, y)

	// int → string via strconv.Itoa
	n := 42
	s := strconv.Itoa(n)
	fmt.Printf("int %d -> string %q\n", n, s)

	// string → int via strconv.Atoi
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("conversion failed:", err)
	}
	fmt.Printf("string %q -> int %d\n", str, num)

	// float → string
	fs := strconv.FormatFloat(3.14, 'f', 2, 64)
	fmt.Println("float as string:", fs)

	// string → float
	pi, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("parsed pi:", pi)

	// byte/rune <-> string
	bs := []byte("hello")
	rs := []rune("नमस्ते")
	fmt.Println(bs, rs)
	fmt.Println(string(bs), string(rs))
}
