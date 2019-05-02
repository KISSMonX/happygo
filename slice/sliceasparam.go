package main

import (
	"fmt"
	"time"
)

func main() {
	sl := []int{1, 2, 3}

	go func(s []int) {
		fmt.Println(s)
	}(sl)

	time.Sleep(time.Second)
}
