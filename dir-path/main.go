package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("/Users/Denzy/example.py")

	fmt.Printf("file= %v\n", file)
	fmt.Printf("dir= %v\n", dir)

	fmt.Println("------------------")

	// more concise version
	dir2, file2 := path.Split("Users/Denzy/example.sql")
	fmt.Printf("file= %v\n", file2)
	fmt.Printf("dir= %v\n", dir2)

}
