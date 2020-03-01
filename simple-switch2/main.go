package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: [Enter any month of the year]")
		return
	}

	switch month := strings.ToLower(os.Args[1]); month {
	case "december", "january", "february":
		fmt.Println("It is Winter 🥶️️")
	case "march", "april", "may":
		fmt.Println("It is Spring 🌸")
	case "june", "july", "august":
		fmt.Println("It is Summertime 😎")
	case "september", "october", "november":
		fmt.Println("It is Autumn 🍁")
	default:
		fmt.Printf("%q is not a valid entry.\n", month)
	}
}
