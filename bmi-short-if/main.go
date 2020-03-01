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
	var (
		h, w float64
		err  error
	)
	/*
		When initializing short if-statements, the simple statement (the short-declaration) comes first.
		Then the conditional expression comes after the semicolon.
	*/

	if arg := os.Args; len(arg) != 3 {
		fmt.Println(`Usage: Please enter your "Height" and "Weight"...`)
	} else if h, err = strconv.ParseFloat(arg[1], 64); err != nil {
		fmt.Printf("Invalid input: %q\n", arg[1])
	} else if w, err = strconv.ParseFloat(arg[2], 64); err != nil {
		fmt.Printf("Invalid input: %q\n", arg[2])
	} else {
		bmi := w / (h * h)
		fmt.Println(strings.TrimSpace(msg))
		fmt.Println()
		fmt.Printf("Weight: %.1fkg, Height: %.2fm, Body Mass Index: %.2f\n", w, h, bmi)
	}

}
