package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const msg = `
BMI commandline App
-------------------
This app computes your Body Mass Index using your "Height" and "Weight" values.
`

	if len(os.Args) < 3 {
		fmt.Println(`Please enter your "Height" and "Weight"...`)
	} else {
		var (
			h    float64
			w    float64
			arg1 = os.Args[1]
			arg2 = os.Args[2]
		)
		h, _ = strconv.ParseFloat(arg1, 64)
		w, _ = strconv.ParseFloat(arg2, 64)
		bmi := w / (h * h) // weight in kg, height in meters

		fmt.Println(strings.TrimSpace(msg))
		fmt.Println()
		fmt.Printf("Weight: %.1fkg, Height: %.2fm, Body Mass Index: %.2f\n", w, h, bmi)
	}

}
