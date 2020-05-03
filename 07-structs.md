# 07 Structs

[[Back]](README.md)

In Go a struct is a collection of fields.

```go
package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) toString() string {
	return fmt.Sprintf("X = %d, Y = %d", p.X, p.Y)
}

type NamedPoint struct {
	Point
	Name string
}

type ToString interface {
	toString() string
}

type Cat struct {
	Name string
}

func (c Cat) toString() string {
	return fmt.Sprintf("%s says 'meow'", c.Name)
}

func printAsString(object ToString) {
	fmt.Println(object.toString())
}

func main() {
	myPoint := Point{X: 1, Y: 2}
	printAsString(myPoint)
	myNamedPoint := NamedPoint{Point: Point{X: 3, Y: 4}, Name: "Jeff"}
	printAsString(myNamedPoint)
	cat := Cat{Name: "Kitty"}
	printAsString(cat)
}
```

## Fields

```go
type Point struct {
	X int
	Y int
}
```

```go
myPoint := Point{X: 1, Y: 2}
```

## Methods

```go
func (p Point) toString() string {
	return fmt.Sprintf("X = %d, Y = %d", p.X, p.Y)
}
```

## Embedded Types

```go
type NamedPoint struct {
	Point
	Name string
}
```

```go
myNamedPoint := NamedPoint{Point: Point{X: 3, Y: 4}, Name: "Jeff"}
```

## Interfaces

```go
type ToString interface {
	toString() string
}
```

```go
func printAsString(object ToString) {
	fmt.Println(object.toString())
}
```

[[Back]](README.md)
