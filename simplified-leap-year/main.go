package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		year int64
		err  error
	)

	if arg := os.Args; len(arg) != 2 {
		fmt.Println("Enter year: [year]")
	} else if year, err = strconv.ParseInt(arg[1], 10, 64); err != nil {
		fmt.Printf("Invalid input: %q\n", arg[1])
	} else if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
