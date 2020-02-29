package main

import "fmt"

/*
- all basic literals are typeless --> typeless constant values
- an untyped numeric constant can be assigned any numeric value.
*/

func main(){
	const (
		// typeless constants
		min = 1
		pi = 3.14
		model = "3.7.0" + "-alpha"
		debug = !true
	)
	fmt.Printf("%T\t %T\t %T\t\t %T\t\n", min, pi, model, debug)
	fmt.Printf("%v\t %v\t\t %v\t %v\t\n", min, pi, model, debug)

	const val = 77
	var (
		i int = val
		f float64 = val
		b byte = val
		r rune =val
	)

	fmt.Printf("i : %T f : %T b : %T r : %T\n", i, f, b, r)
	fmt.Printf("i : %d f : %2f b : %b r : %d\n", i, f, b, r)

}
