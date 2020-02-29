package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	age := os.Args[1]

	num, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}
	fmt.Printf("Success: converted %q to %d. \n", age, num)
}
