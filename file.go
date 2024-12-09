package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("ms.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	f, error := os.Open("ms.txt")
	defer f.Close()

	if error != nil {
		fmt.Println(error)

	}

	reader := bufio.NewReader(f)
	b1, error := reader.Peek(10)

	if error != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b1))
}
