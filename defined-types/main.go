package main

import "fmt"

/*
You define own types for two main reasons:
1. Type safety
2. Readability

A defined type inherits its underlying type's operations and methods.
*/

type gram float64
type ounce float64

func main() {
	var g gram = 1000
	var o ounce
	o = ounce(g) * 0.035274
	fmt.Printf("%g grams is equivalent to %.2f ounces.\n", g, o)

}
