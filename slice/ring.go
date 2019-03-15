package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(100)

	for i := 0; i < 130; i++ {
		r.Value = i
		r = r.Next()
	}

	r = r.Move(0)
	fmt.Println(r.Len(), r)
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
}
