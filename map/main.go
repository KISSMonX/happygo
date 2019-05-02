package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	f := func(key, value interface{}) bool {
		fmt.Println("遍历", key, value)
		return true
	}

	m.Store("侯名", "666")
	m.Range(f)

	m.Store("houming", "888")
	m.Range(f)

	m.Store("侯名", "16888")
	m.Range(f)
}
