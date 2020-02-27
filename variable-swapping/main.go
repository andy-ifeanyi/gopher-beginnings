package main

import "fmt"

func main() {
	var (
		currentAmount = 100
		prevAmount    = 50
	)

	currentAmount, prevAmount = prevAmount, currentAmount
	fmt.Println(currentAmount / prevAmount)
}
