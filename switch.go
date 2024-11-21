package main

import "fmt"

func main() {
	// fmt.Printf("vale")

	value := 3
	test := 10

	switch value {
	case 1:
		fmt.Printf("One")
	case 2:
		fmt.Printf("Two")
	case 3:
		fmt.Printf("Three")
	case 4:
		fmt.Printf("Four")
	}

	for a := 0; a <= 10; a++ {
		fmt.Printf("Hi Business Man Msksurjith \n")
	}

	m := 5
	for m < 10 {
		fmt.Printf("How are you Msksurjith \n")
		m++
	}
	fmt.Printf("\nProgram finished")

}
