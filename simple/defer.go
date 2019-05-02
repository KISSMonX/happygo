package main

import "fmt"

func main() {
	fmt.Println(DeferFunc2(1))
}

func DeferFunc2(i int) int {
	t := i

	defer func() {
		t += 3
	}()

	return t
}
