package main

import (
	"fmt"
	"github.com/andy-ifeanyi/gopher/iota-const/months"
	)


/*
An iota is used to generate serially incrementing numbers for contants.
*/
func main() {
	const (
		sunday = iota + 1
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)
	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday)
	fmt.Println(January, February, March, April, May, June, July, August, September, October, November, December)
}
