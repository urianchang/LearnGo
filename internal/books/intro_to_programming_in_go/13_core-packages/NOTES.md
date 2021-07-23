# The Core Packages

## Strings
Contains large amount of functions for handling strings.

```
// true
strings.Contains("test", "es"),
// 2
strings.Count("test", "t"),
// true
strings.HasPrefix("test", "te"),
// true
strings.HasSuffix("test", "st"),
// 1
strings.Index("test", "e"),
// "a-b"
strings.Join([]string{"a","b"}, "-"),
// == "aaaaa"
strings.Repeat("a", 5),
// "bbaa"
strings.Replace("aaaa", "a", "b", 2),
// []string{"a","b","c","d","e"}
strings.Split("a-b-c-d-e", "-"),
// "test"
strings.ToLower("TEST"),
// "TEST"
strings.ToUpper("test"),
```

To convert a string to a slice of bytes and vice versa:

```
arr := []byte("test")
str := string([]byte{'t','e','s','t'})
```

## Input/Output
The `io` package consists of a few functions, but mostly interfaces used in
other packages. Two main interfaces are `Reader` and `Writer`.

`Copy` function copies data from a `Reader` to a `Writer`:

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

To read or write to a `[]byte` or a `string` you can use the `Buffer` struct.

```
var buf bytes.Buffer
buf.Write([]byte("test"))
```

A `Buffer` doesn't have to be initialized and supports both the `Reader` and
`Writer` interfaces. You can convert it into a `[]byte` by calling
`buf.Bytes()`. If you only need to read from a string you can use
`strings.NewReader` function (which is more efficient).

## Files & Folders
To open a file in Go use the `Open` function from the `os` package. We use
`defer file.Close()` right after opening the file to make sure the file is
closed as soon as the function completes.

To get the contents of a directory, we use the same `os.Open` function but give
it a directory path instead of a file name. Then we call the `Readdir` method.
There's a `Walk` function provided in the `path/filepath` package that
recursively walks a folder. The function you pass to `Walk` is called for
every file and folder in the root folder.

## Errors
Go has a built-in type for errors. We can create our own errors by using the
`New` function in the `errors` package:

```
package main

import "errors"

func main() {
    err := errors.New("error message")
}
```

## Containers & Sort
The `container/list` package implements a doubly-linked list. The zero value for
a `List` is an empty list. Values are appended to the list using `PushBack`. We
loop over each item in the list by getting the first element, and following all
the links until we reach nil.

The sort package contains functions for sorting arbitrary data. The `Sort`
function in `sort` takes a `sort.Interface` and sorts it. The `sort.Interface`
requires 3 methods: `Len`, `Less`, and `Swap`.

## Hashes & Cryptography
A hash function takes a set of data and reduces it to a smaller fixed size.
There are two categories: cryptographic (e.g. `sha1`) and non-cryptographic (
e.g. `crc32`).

## Synchronization Primitives
Preferred way to handle concurrency and synchronization in Go is through
goroutines and channels. However, Go does provide more traditional
multithreading routines in the `sync` and `sync/atomic` packages.

### Mutexes
A mutex (mutual exclusive lock) locks a section of code to a single thread at a
time and is used to protect shared resources from non-atomic operations.

When a mutex is locked, any other attempt to lock it will block until it is
unlocked. Great care should be taken when using mutexes or the synchronization
primitives in the `sync/atomic` package.

One of Go's biggest strengths is that the concurrency features it provides are
much easier to understand and use properly than threads and locks.
