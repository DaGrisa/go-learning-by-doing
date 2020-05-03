# 01 - Hello Go

[[Back]](README.md)

The standard "Hello World" example written in go.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello go!")
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
import "fmt"
```

Other packages can be imported using the `import` keyword.

### Go Standard Library

Available packages in the standard library can be found [here](https://golang.org/pkg/). In this first example we are using the `fmt` package for formatted output ("printing").

[[Back]](README.md)
