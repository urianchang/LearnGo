# Packages
Go provides a mechanism fro code reuse: packages.

3 purposes:
1. Reduces chances of having overlapping names. Keeps function names short and
succinct.
2. It organizes code so that its easier to find code you want to reuse.
3. Speeds up compiler by only requiring recompilation of smaller chunks of a
program.

In Go, if something starts with a capital letter that means other packages
(and programs) are able to see it. If we had named the function `average`
instead of `Average`, our `main` program would not have been able to see it.

**NOTE**: Package names match the folders they fall in.

