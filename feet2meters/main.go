package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	type (
		Feet   float64
		Meters float64
	)

	const (
		feet2meter = 0.3048
		msg        = `
Convert Feet to Meters
----------------------
This commandline app converts feet to meters.`
	)
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument in order to run this program.")
	} else {
		var (
			feet   Feet
			meters Meters
			arg    = os.Args[1]
		)
		res, _ := strconv.ParseFloat(arg, 64)
		feet = Feet(res)

		meters = Meters(feet * feet2meter)

		fmt.Println(strings.TrimSpace(msg))
		fmt.Printf("%v feet --> %.2f meters.\n", feet, meters)
	}

}
