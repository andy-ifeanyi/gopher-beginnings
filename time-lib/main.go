package main

import (
	"fmt"
	"time"
)

func main() {
	h, _ := time.ParseDuration("3h30m")

	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of h's underlying type: %T\n", int64(h))

	fmt.Println(h.Hours(), "hours")
	fmt.Println(int64(h))

	var m int64 = 2

	h *= time.Duration(m)

	fmt.Printf("%v\n", h)
}
