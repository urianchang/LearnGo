# Arrays, Slices, and Maps

## Arrays
An array is a numbered sequence of elements of a single type with a fixed
length.

```
var x [5]int
```

Shorter syntax for creating arrays:

```
x := [5]float64{ 98, 93, 77, 82, 83 }
```

## Slices
A slice is a segment of an array. Like arrays slices are indexable and have a
length. Unlike arrays this length is allowed to change.

```
var x []float64
```

To create a slice, use the `make` function:

```
x := make([]float64, 5)
```

`append` creates a new slice by taking an existing slice and appending all the
arguments to it.

```
func main() {
    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
}
```

`copy` copies contents from one slice into another.

```
func main() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}
```

## Maps
A map is an unordered collection of key-value pairs. They have to be initialized
before they can be used.

```
x := make(map[string]int)
x["key"] = 10
fmt.Println(x["key"])
```

`delete` can be used to delete items from a map.

```
delete(x, 1)
```







