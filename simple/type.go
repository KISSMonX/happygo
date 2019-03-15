package main

import "fmt"

func False() bool {
	return false
}

func main() {
	switch False() {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}
