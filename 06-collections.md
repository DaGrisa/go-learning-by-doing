# Collections

[[Back]](README.md)

## Array

Is a numbered sequence of a single type with a fixed length.

```go
var array [5]int
```

## Slice

Is a segment of an array, the length is allowed to change.

```go
var slice []int
```

Go supports slices with two built-in functions `append` and `copy`.

```go
var sliceOne []int
var sliceTwo []int
sliceOne = append(sliceOne, 1)
sliceOne = append(sliceOne, 2)
sliceTwo = append(sliceTwo, 3)
copy(sliceOne, sliceTwo)
```

## Map

Is an unordered collection of key-value pairs.

```go
areaCode := make(map[string]int)
areaCode["AT"] = 43
areaCode["CH"] = 41
areaCode["DE"] = 49
areaCode["XX"] = 99
delete(areaCode, "XX")
areaCodeXX, exists := areaCode["XX"]
```

The built-in function `delete` provides deletion of map elements.

```go
delete(areaCode, "XX")
```

Accessing an element of a map returns 2 values: the first is the value, the second one a boolean that shows, if the selected key exists in the map.

```go
areaCodeXX, exists := areaCode["XX"]
```

[[Back]](README.md)
