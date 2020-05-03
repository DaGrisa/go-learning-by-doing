# 02 Variables

[[Back]](README.md)

Usage examples of variables in go.

```go
package main

import "fmt"

var packageLevel int = 42

func main() {
	var functionLevel, otherVariable bool = true, true
	implicitType, otherInteger := "wow", 24

	fmt.Println(
		packageLevel,
		functionLevel,
		otherVariable,
		implicitType,
		otherInteger)
}
```

## Declaration

Variable in go can be declared at package and function level.

```go
package main

var packageLevel int = 42

func main() {
    var functionLevel, otherVariable bool = true, true
    ...
}
```

## Definition

Variables can be defined at declaration by using `=`.

```go
var packageLevel int = 42
```

### Implicit Types

Variables can be implicitly typed by using `:=`.

```go
	implicitType, otherInteger := "wow", 24
```

[[Back]](README.md)
