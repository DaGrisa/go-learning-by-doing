# 01 - Hello Go

[[Back]](README.md)

The standard "Hello World" example written in Go.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello go! Here you have some", math.Pi)
}
```

## Program Entry

```go
func main() {
	...
}
```

## Package Declaration

```go
package main
```

Every go program is made out of packages and start running in package `main`.

## Importing Packages

```go
import (
	"fmt"
	"math"
)
```

Other packages can be imported using the `import` keyword.

### Go Standard Library

Available packages in the standard library can be found [here](https://golang.org/pkg/). In this first example we are using the `fmt` and `math` package.

### Exported Names

Names beginning with a capital letter are exported by Go. Only exported names can be referred from imported packages. In this expample it is the exported variable Pi with `math.Pi`.

[[Back]](README.md)
