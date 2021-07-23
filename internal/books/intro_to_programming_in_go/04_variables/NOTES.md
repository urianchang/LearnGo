# Variables
Variables in Go are created by first using the `var` keyword, then specifying
the variable name, the type, and finally assigning a value to the variable.

```
var x string
x = "Hello World"
```

Shorter statement that uses the Go compiler to infer the type based on the
literal value you assign the variable.

```
x := "Hello World"
// OR
var x = "Hello World"
```

Use lower camel case (e.g. "dogsName") to represent multiple words in a variable
name.

**Scope:** Go is lexically scoped using blocks.

Constants are basically variables whose values cannot be changed later. They are
created in the same way you create variables but instead of using `var`, use
`const`.

## Defining Multiple Variables
Use the keyword (i.e. `var` or `const`) followed by parentheses with each
variable on its own line.

```
var (
    a = 5
    b = 10
    c = 15
)
```
