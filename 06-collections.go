package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Integer array:", array)

	var sliceOne []int
	var sliceTwo []int
	sliceOne = append(sliceOne, 1)
	sliceOne = append(sliceOne, 2)
	sliceTwo = append(sliceTwo, 3)
	copy(sliceOne, sliceTwo)
	fmt.Println("Slice after copy:", sliceOne)

	areaCode := make(map[string]int)
	areaCode["AT"] = 43
	areaCode["CH"] = 41
	areaCode["DE"] = 49
	areaCode["XX"] = 99
	fmt.Println("Area code map:", areaCode)
	delete(areaCode, "XX")
	areaCodeXX, exists := areaCode["XX"]
	fmt.Println("Area code XX exists:", exists, areaCodeXX)
}
