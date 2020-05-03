# 04 - Basic Types

```go
package main

import "fmt"

func returnDefaults() (bool, int, string) {
	var boolean bool
	var integer int
	var charSequence string
	return boolean, integer, charSequence
}

func main() {
	boolean, integer, charSequence := returnDefaults()
	fmt.Println("bool, int and string default to", boolean, integer, charSequence)

	var maxSizeInt32 int32 = 0
	for i := int32(0); i >= 0; i++ {
		maxSizeInt32 = i
	}
	fmt.Println("Max size of int32 =", maxSizeInt32)

	var minSizeInt32 int32 = 0
	for i := int32(0); i <= 0; i-- {
		minSizeInt32 = i
	}
	fmt.Println("Min size of int32 =", minSizeInt32)

	integerToConvert := 42
	floatResult := float64(integerToConvert)
	unsignedResult := uint(integerToConvert)
	fmt.Printf("%d after conversion has type %T and %T", integerToConvert, floatResult, unsignedResult)
}
```

## Boolean

Booleans are represented by the type `bool` and default to `false`.

## String

Character sequences are declared with the type `string` and default to `""`.

## Numbers

### Integer

Integers are available as "signed" and "unsigned" (only positive numbers) and in different sizes.

```go
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
```

The size of `int` and `uint` is platform dependent.
`uintptr` has the length of a pointer and should not be used (more detials [here](https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang)).

### Float

Floating point numbers are declared as `float32` and `float64`.

### Complex

Data types for complex numbers are `complex64` and `complex128`.

### Numeric Constants

Numeric constants on Go are high precision values and therefor mathematically exact.

## Alias

### byte

`byte` is an alias for int8.

### rune

`rune` represents a Unicode code point and is an alias for int32.

## Type Conversion

There is no implicit type conversion in Go, only explicit.

```go
	integerToConvert := 42
	floatResult := float64(integerToConvert)
	unsignedResult := uint(integerToConvert)
```
