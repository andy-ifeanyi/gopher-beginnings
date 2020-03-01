package main

import (
	"fmt"
	"os"
	"strconv"
)

const feet2meter = 0.3048

func main() {

	if arg := os.Args; len(arg) != 2 {
		fmt.Println("usage: [ in meters]")
	} else if f, err := strconv.ParseFloat(arg[1], 64); err != nil {
		fmt.Printf("Cannot convert %\n", f)
	} else {
		m := f * feet2meter
		fmt.Printf("%v feet --> %.2f meters\n", f, m)
	}
}
