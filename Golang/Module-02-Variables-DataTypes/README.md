# Module 02 - Variables, Constants & Data Types

## Variables

Go is **statically typed** - every variable has a type fixed at compile time.

### Three ways to declare:

```go
var x int = 10            // explicit type + value
var y = 20                // type inferred
z := 30                   // short form (only inside functions)
```

### Zero values
Uninitialized variables get a **zero value**:
- numeric → `0`
- string → `""`
- bool → `false`
- pointer/slice/map/chan/func/interface → `nil`

## Constants

Declared with `const`. Value fixed at compile time, cannot change.

```go
const Pi = 3.14
const (
    StatusOK    = 200
    StatusNotFound = 404
)
```

### `iota`
Auto-incrementing constant generator - great for enums.

```go
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
)
```

## Primitive Types

| Category | Types |
|----------|-------|
| Integer  | `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8` … |
| Float    | `float32`, `float64` |
| Complex  | `complex64`, `complex128` |
| Text     | `string`, `rune` (alias for `int32`), `byte` (alias for `uint8`) |
| Boolean  | `bool` |

## Type Conversion

Go requires **explicit** conversions - no implicit casting.

```go
i := 42
f := float64(i)
s := strconv.Itoa(i) // int → string
```

## Files in this module

- [01_variables.go](01_variables.go)
- [02_constants.go](02_constants.go)
- [03_data_types.go](03_data_types.go)
- [04_type_conversion.go](04_type_conversion.go)
- [05_iota_enum.go](05_iota_enum.go)

## Exercises
1. Declare variables using all three declaration styles.
2. Create constants for HTTP status codes (200, 404, 500) using `iota`-style groups.
3. Convert a float to int and note what happens to the decimal part.
4. Read a number as a string and convert it to int using `strconv.Atoi`.
