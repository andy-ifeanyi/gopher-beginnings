package main

import (
	"fmt"
	"time"
)

/*
In Golang, a switch statement is actually a syntactic-sugar for an if statement.
*/
func main() {
	switch t := time.Now().Hour(); {
	case t > 21:
		fmt.Println("Night")
	case t > 17:
		fmt.Println("Evening")
	case t > 12:
		fmt.Println("Afternoon")
	case t <= 11:
		fmt.Println("Morning")
	}
}
