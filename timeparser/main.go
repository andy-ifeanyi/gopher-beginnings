package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	t, _ := time.ParseDuration("7h45m")

	arg := os.Args[1]

	multiplier, _ := strconv.ParseInt(arg, 10, 64)

	t *= time.Duration(multiplier)

	fmt.Println(t)

}
