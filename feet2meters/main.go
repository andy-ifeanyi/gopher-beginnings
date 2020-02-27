package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%v feet --> %v meters.\n", feet, meters)
}