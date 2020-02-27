package main

import "fmt"

func main() {
	// fmt.Printf("%b\n", 4)
	// fmt.Printf("%02b\n", 4)
	// fmt.Printf("%04b\n", 4)
	// fmt.Printf("%08b\n", 4)

	// for i := 10000000; i < 10000100; i++ {
	// 	fmt.Printf("%d \t %b \t %x \n", i, i, i)
	// }

	var b byte
	b = 0
	fmt.Printf("Minimum byte value: %08b = %d\n", b, b)
	b = 255
	fmt.Printf("Maximum byte value: %08b = %d\n", b, b)
	fmt.Printf("Maximum byte value: %08b = %d\n", 2, 2)
}
