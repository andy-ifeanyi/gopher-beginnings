package main

import (
	"fmt"
	"os"
)

func main(){
	// fmt.Printf("%#v\n", os.Args)

	// fmt.Println("Path length:", len(os.Args))

	// fmt.Println("Path:", os.Args[0])
	// fmt.Println("1st arg:", os.Args[1])
	// fmt.Println("2nd arg:", os.Args[2])

	name := os.Args[1]

	fmt.Println("Hello courageous", name, "!")

	name, birth := "Denzy", 1980
	
	fmt.Println("My name is", name)
	fmt.Println("I came to this planet in", birth)
    fmt.Println("Thou art blessed oh Denzy!")

}