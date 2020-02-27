package main

import (
	"fmt"
)

/*
A defined type and its soure type share the same underlying type.
*/

func main() {
	type (
		// Gram's underlying type is int
		Gram int

		// Kilogram's underlying type is int
		Kilogram Gram

		// Ton's underlying type is int
		Ton Kilogram
	)

	var (
		salt  Gram     = 100
		rice  Kilogram = 5
		truck Ton      = 10
	)

	fmt.Printf("salt: %d, rice: %d, truck: %d\n", salt, rice, truck)

	// salt = rice // will raise underlying type error.
	salt = Gram(rice)
	rice = Kilogram(truck)
	//truck = Ton(Gram(int(rice)))
	truck = Ton(rice)
	fmt.Printf("salt: %d, rice: %d, truck: %d\n", salt, rice, truck)

}
