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

	// mixed variable declaration

	var a, b, c = 10, 12, "msksurjith"

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%s\n", c)

	// Desicion making statement
	m := 10

	if m == 5 {
		fmt.Printf("A is equal to 10")
	} else {
		fmt.Printf("A is not equal to 10 \n")
	}

	l := 5
	n := 7
	o := 8
var data string := "hello"
var num :=[5]int{10,20,30,40,50}
	for i:=0;i<5;i++{
		fmt.Printf("%d\n",num[i])
	}
	
	if l > n {
		if n > o {
			fmt.Printf("L is Big \n")
		} else {
			fmt.Print("O is Big \n")
		}
	}
	ans := sum(10, 20)
	fmt.Printf("answer is %d \n", ans)
	// var news := "welcome"
	
	// array value

	
}
func sum(a, b int) int {
	c := a + b
	return c

	
}
