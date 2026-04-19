// Run: go run 05_iota_enum.go
//
// `iota` is Go's way to create auto-incrementing constants — perfect for enums.
// It starts at 0 inside a const block and increments by 1 per line.

package main

import "fmt"

// Days of the week
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
	Wednesday     // 3
	Thursday      // 4
	Friday        // 5
	Saturday      // 6
)

// Bit flags — using iota with bit shifting.
const (
	Read    = 1 << iota // 1
	Write               // 2
	Execute             // 4
)

// Skip values with _
const (
	_  = iota // ignore 0
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Println("Monday:", Monday, "Friday:", Friday)
	fmt.Println("Permissions: R=", Read, "W=", Write, "X=", Execute)

	perm := Read | Write
	fmt.Printf("Combined R|W = %b (%d)\n", perm, perm)

	fmt.Println("KB:", KB, "MB:", MB, "GB:", GB, "TB:", TB)
}
