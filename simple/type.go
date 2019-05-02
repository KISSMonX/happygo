package main

import "fmt"

type W struct {
	Id int
	t  *T
}

func createT() *T {
	var t *T = &T{}
	t.Id = 1
	return t
}

func (w *W) setT() {
	t := createT()
	fmt.Println(t.Id)
	fmt.Println("setT t地址: ", &t)
	w.t = t
	fmt.Println("setT t地址: ", &w.t)
}

type T struct {
	Id   int
	Name string
}

func main() {
	w := W{}
	w.setT()

	fmt.Println(&w.t)
}
