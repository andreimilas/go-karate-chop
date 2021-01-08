# CodeKata - Kata02: Karate Chop

## Author
Andrei Milas

## Description
Proposed solutions for Kata02 in Go (http://codekata.com/kata/kata02-karate-chop/).
The project contains 5 implementations.

### builtin 
* It uses the `sort` standard library package
* It's simple, efficient and most likely to use in a production environment
### iterative
* The checks are performed in a loop and the relevant dataset is halved at each iteration
### functional
* A loop is used, as for the iterative approach.
* The checks are done inside a function, taking advantage of Go's multiple return values support.
* The index and slice are changed at each iteration (instead of just the start / end marks in the iterative version).
### recursive
* A halving function is called recursively instead of a loop.
* The recursion stops when a length 1 slice is reached.
### structural
* An iterative approach using a Go structure as an input.
* A method is used instead of a function with the aforementioned structure as a receiver.
* It's more complex than the standard iterative approach and requires initializing the receiver with the given input slice.
