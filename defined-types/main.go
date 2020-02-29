package main

import "fmt"

/*
You define own types for two main reasons:
1. Type safety
2. Readability
3. To declare new methods

Note:
 - methods can only be declared on a defined types.
 - a defined type doesn't get its source type's methods.

A defined type only gets its underlying type's:
- representation
- size
- range of values
*/

type gram float64
type ounce float64

func main() {
	var g gram = 1000
	var o ounce
	o = ounce(g) * 0.035274
	fmt.Printf("%g grams is equivalent to %.2f ounces.\n", g, o)

}
