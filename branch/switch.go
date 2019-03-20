package main

import "fmt"

func main() {
	// 执行汇编是否有区别
	fswitch1(1)
	fswitch2(1)
}

func fswitch1(x int) {
	switch x {
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("default")
	}
}

func fswitch2(x int) {
	switch x {
	default:
		fmt.Println("default")
	case 1:
		fmt.Println("1")
	}
}
