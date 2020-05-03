package main

import "fmt"

func fibonacci(n int) int {
	if(n == 0) {
		return 0
	}
	if(n == 1) {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2);
}

func printFibonacci(n int) {
    fmt.Printf("Value of index %d in the fibonacci sequence is %d.\n", n, fibonacci(n))
}

func main() {
	n := 42
    printFibonacci(n)
}
