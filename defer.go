package main

import "fmt"

func main() {
	defer PrintText("defered")
	defer PrintText("second")
	PrintText("Hellow")
	defer PrintText("last")
}
func PrintText(p string) {
	fmt.Println(p)
}
