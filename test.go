package main

import "fmt"

const max = 3

func main() {
	var a int = 10
	var ptr *int
	ptr = &a
	fmt.Printf("Address of a variable is %d \n", &a)
	fmt.Printf("Address store in ptr is %x \n", *ptr)
	fmt.Printf("Address of *ptr %d \n", *ptr)

	var nptr *int
	fmt.Printf("Null pointer is %x", nptr)

	arr := []int{10, 20, 30, 40, 50}
	var i int
	var p [max]*int
	for i = 0; i < max; i++ {
		p[i] = &arr[i]
	}
	for i = 0; i < max; i++ {
		fmt.Printf("Value of a[%d] = %d \n", i, *p[i])
	}

	var x int
	var ptr1 *int
	var pptr **int

	x = 3000

	ptr1 = &x
	pptr = &ptr1

	fmt.Printf("Value of a = %d\n", x)
	fmt.Printf("Value available at *ptr = %d\n", *ptr1)
	fmt.Printf("Value available at **ptr %d\n", **pptr)
}
