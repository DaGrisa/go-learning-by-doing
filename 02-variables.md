# 02 - Variables

[[Back]](README.md)

Usage examples of variables in Go.

```go
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
```

## Declaration

Variables in Go can be declared at package and function level.

```go
package main

var packageLevel int = 42

func main() {
    var functionLevel, otherVariable bool = true, true
    ...
}
```

### Default Values

Variables can be defaulted at declaration by using `=`.

```go
var packageLevel int = 42
```

### Implicit Types

Variables can be implicitly typed by using `:=` (only available for function-level variables).

```go
	implicitType, otherInteger := "wow", 24
```

### Constants

Constants can be defined using the keyword `const`. They can be character, string, boolean, or numeric values and not be declared using `:=`.

```go
	const constant = 42
```

[[Back]](README.md)
