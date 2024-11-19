package main

import "fmt"

func main() {
	fmt.Printf("hello, businessman msksurjith\n")
	// static declaration variable
	var num1 int
	num1 = 100
	fmt.Printf("%d\n", num1)

	var floatvalue float32
	floatvalue = 100.1
	fmt.Printf("Value of num1 is: %d\n", num1)
	fmt.Printf("Value of Float is: %f\n", floatvalue)

	// dynamic declaration variable

	x := 1000
	y := 10.22
	name := "msk"
	fmt.Printf("%d\n", x)
	fmt.Printf("%f\n", y)
	fmt.Printf("%s\n", name)

}
