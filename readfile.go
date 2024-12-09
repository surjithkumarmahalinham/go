package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	message := []byte("Hello World")
	err := ioutil.WriteFile("ms.txt", message, 0644)

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("newms.txt")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("Hello new file created !!")
	f.Sync()

	w := bufio.NewWriter(f)

	w.WriteString(" by msksurjith")

	w.Flush()

}
