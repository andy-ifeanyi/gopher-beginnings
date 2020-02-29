package main

import (
	"fmt"

	"github.com/andy-ifeanyi/gopher-beginnings/underlying-types/measures"
)

/*
A defined type and its source type share the same underlying type.
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
		spice Gram     = 100
		rice  Kilogram = 5
		truck Ton      = 10
	)

	fmt.Printf("spice: %d, rice: %d, truck: %d\n", spice, rice, truck)

	// salt = rice // will raise underlying type error.
	spice = Gram(rice)
	rice = Kilogram(truck)
	//truck = Ton(Gram(int(rice)))
	truck = Ton(rice)
	fmt.Printf("spice: %d, rice: %d, truck: %d\n", spice, rice, truck)

	type iGram measures.Gram
	var milk iGram = 40

	milk = iGram(spice)
	fmt.Printf("milk: %d, rice: %d, truck: %d\n", milk, rice, truck)

}
