# Structs and Interfaces
A struct is a type which contains named fields. For example, we could represent
a circle like this:

```
type Circle struct {
    x float64
    y float64
    r float64
}
```

The `type` keyword introduces a new type. It's followed by the name of the type.
The keyword `struct` indicates that a `struct` type is being defined. Each field
has a name and a type. We can collapse fields that have the same type:

```
type Circle struct {
    x, y, r float64
}
```

We can create an instance of a struct in a few ways:

```
// local variable
var c Circle

// allocates memory for all fields, sets each of them to zero value and returns
// a pointer (*Circle)
c := new(Circle)

// initialize with values
c := Circle{x: 0, y: 0, r: 5}
// OR
c := Circle{0, 0, 5}
```

We can access fields using the `.` operator.

```
fmt.Println(c.x, c.y, c.r)
c.x = 10
c.y = 5
```

To create a method, we add a "receiver" between `func` and the name of the
function. A receiver has a name and a type.

```
func (c *Circle) area() float64 {
     return math.Pi * c.r*c.r
}
```

## Embedded Types
A struct's fields usually represent the has-a relationship. Also known as
anonymous fields, embedded types look like this:

```
type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
type Android struct {
    Person
    Model string
}
```

When defined this way, the `Person` struct can be accessed using the type name:

```
a := new(Android)
a.Person.Talk()
```

We can also call any `Person` methods directly on `Android`:

```
a := new(Android)
a.Talk()
```

## Interfaces

```
type Shape interface {
     area() float64
}
```

Like a struct, an interface is created using the `type` keyword, followed by a
name and the keyword `interface`. But instead of defining fields, we define a
"method set". A method set is a list of methods that a type must have in order
to implement the interface.

We can use interface types as arguments to functions:

```
func totalArea(shapes ...Shape) float64 {
     var area float64
     for _, s := range shapes {
           area += s.area()
     }
     return area
}

fmt.Println(totalArea(&c, &r))
```

Interfaces can also be used as fields:

```
type MultiShape struct {
     shapes []Shape
}
```

We can change `MultiShape` into a `Shape` by giving it an area method:

```
func (m *MultiShape) area() float64 {
     var area float64
     for _, s := range m.shapes {
           area += s.area()
     }
     return area
}
```

Now `MultiShape` can contain `Circle`, `Rectangle`, or other `MultiShape`.
