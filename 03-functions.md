# 03 - Functions

[[Back]](README.md)

The recursive fibonnaci sequence as a function example in Go.

```go
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
```

## Function Declaration

A function can have parameters (0-n) and return types (0-n).

```go
func fibonacci(n int) int {
	...
}

func printFibonacci(n int) {
    fmt.Printf("Value of index %d in the fibonacci sequence is %d.\n", n, fibonacci(n))
}
```

```go
func swap(a, b string) (string, string) {
    return b, a
}
```

You can also declare return variables in function header.

```go
func splitDateTime(string isoDate) (date, time, offset string) {
    dateTime = strings.Split(isoDate, "T")
    date = dateTime[0]
    time = dateTime[1]
	return
}
```

[[Back]](README.md)
