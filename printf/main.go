package main

import "fmt"

func main() {
	// fmt.Printf("Decimal\t\t Binary\t Hexadecimal\t Octal\n")
	// for i := 50; i < 100; i++ {
	// 	fmt.Printf("The Number %v:\t %b\t %x\t\t %o\t\n", i, i, i, i)

	var (
		parasite         = "Malaria Falciparum"
		incubationPeriod = 7
		isTreatable      = true
	)
	fmt.Printf("Malaria is caused by %s. It has an incubation period of %d days to 4 weeks. Treatable: %t.\n", parasite, incubationPeriod, isTreatable)

	celsius := 35.0
	fahrenheit := (9*celsius + 160) / 5
	fmt.Printf("%g°C == %g°F\n", celsius, fahrenheit)
}
