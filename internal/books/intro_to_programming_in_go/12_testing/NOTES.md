# Testing
The `go test` command will look for any tests in any of the files in the
current folder and run them. Tests are identified by starting a function with
the word `Test` and taking one argument of type `*testing.T`.

A very common way to setup tests is to create a `struct` to represent the inputs
and outputs for the function. Then create a list of the structs (pairs). Next we
loop through each one and run the function.

Refer to the `math_test.go` file in Ch 11.
