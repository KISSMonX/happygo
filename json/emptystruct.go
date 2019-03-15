package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Model struct {
		A string `json:"A"`
		B bool   `json:"B"`
	}

	var xMod []Model
	b, err := json.Marshal(xMod)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	yMod := []Model{}

	b, err = json.Marshal(yMod)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

// output:
// null
// []
