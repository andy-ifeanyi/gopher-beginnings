package main

import "fmt"

func main() {
	var counter int
	counter++
	counter++
	counter++
	counter--
	fmt.Printf("Count = %v\n", counter)
	num1 := 5
	num1 += 2 // assignment operation
	num1++
	fmt.Printf("First number = %v\n", num1)
	num2 := 7
	num2 -= 2 // assignment operation
	num2--
	fmt.Printf("Second number = %v\n", num2)
}
