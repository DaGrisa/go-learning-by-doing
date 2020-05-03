# 05 Control Structures

[[Back]](README.md)

```go
package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	for d := int(2); d < n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func main() {
	month := int(time.Now().Month())
	var season string
	if month >= 3 && month < 6 {
		season = "spring"
	} else if month >= 6 && month < 9 {
		season = "summer"
	} else if month >= 9 && month < 12 {
		season = "fall"
	} else {
		season = "winter"
	}
	fmt.Println("Season is", season)

	hour := time.Now().Hour()
	var daytime string
	switch {
	case hour >= 23:
		daytime = "midnight"
	case hour >= 18:
		daytime = "evening"
	case hour >= 13:
		daytime = "afternoon"
	case hour >= 11:
		daytime = "midday"
	default:
		daytime = "morning"
	}
	fmt.Println("Daytime is", daytime)

	primesUpTo := 250
	var primes []int
	for n := int(1); n <= primesUpTo; n++ {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	fmt.Println("Primes up to", primesUpTo, "=", primes)
}
```

## Conditional Statements

### If-Then-Else

```go
	if month >= 3 && month < 6 {
		season = "spring"
	} else if month >= 6 && month < 9 {
		season = "summer"
	} else if month >= 9 && month < 12 {
		season = "fall"
	} else {
		season = "winter"
	}
```

### Switch-Case

```go
	switch {
	case hour >= 23:
		daytime = "midnight"
	case hour >= 18:
		daytime = "evening"
	case hour >= 13:
		daytime = "afternoon"
	case hour >= 11:
		daytime = "midday"
	default:
		daytime = "morning"
	}
```

```go
	switch hour {
	case 23, 00:
		daytime = "midnight"
	case 18, 19, 20, 21, 22:
		daytime = "evening"
	case 13, 14, 15, 16, 17:
		daytime = "afternoon"
	case 11, 12:
		daytime = "midday"
	default:
		daytime = "morning"
	}
```

## Loops

Go only knows one loop construct, the for-loop.

### For-Loop

```go
func isPrime(n int) bool {
	for d := int(2); d < n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}
```

[[Back]](README.md)
