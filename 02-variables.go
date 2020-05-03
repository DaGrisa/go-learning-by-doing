package main

import "fmt"

var packageLevel int = 42

func main() {
	var functionLevel, otherVariable bool = true, true
	implicitType, otherInteger := "wow", 24
	const constant = 42

	fmt.Println(
		packageLevel,
		functionLevel,
		otherVariable,
		implicitType,
		otherInteger,
		constant)
}
