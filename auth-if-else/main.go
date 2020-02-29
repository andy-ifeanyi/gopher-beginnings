package main

import (
	"fmt"
	"os"
)

const (
	usr1, usr2 = "Admin", "Dev"
	pwd1, pwd2 = "secret1", "secret2"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username][password]")
		return
	}

	u, p := args[1], args[2]

	if u != usr1 && u != usr2 {
		fmt.Println("Invalid Username")
	} else if u == usr1 && p == pwd1 {
		fmt.Printf("%s logged in successfully.\n", u)
	} else if u == usr2 && p == pwd2 {
		fmt.Printf("%s logged in successfully.\n", u)
	} else {
		fmt.Printf("Invalid password for %s.\n", u)
	}
}
