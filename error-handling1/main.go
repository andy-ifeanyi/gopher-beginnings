package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	args := os.Args


	if len(args) !=2  {
		fmt.Println("Usage: [age]")
		return
	}
	age := args[1]

	num, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}
	fmt.Printf("Success: converted %q to %d. \n", age, num)
}
