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
