package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name string `json:"name"` //default key set
		Age  int
		Car  []string
	}

	p1 := &Person{Name: "Msk", Age: 28, Car: []string{"RR", "Tata", "Tesla"}}

	data, _ := json.Marshal(p1)
	fmt.Println(string(data))

	data1 := `{"Name":"Msk","Age":28,"Car":["RR","Tata","Tesla"]}`
	p2 := &Person{}

	json.Unmarshal([]byte(data1), p2)

	fmt.Println(p2.Name, p2.Age)
}
