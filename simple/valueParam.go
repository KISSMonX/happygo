package main

import "fmt"

func main() {
	persons := make(map[string]int)
	persons["侯名"] = 233

	mp := &persons

	fmt.Printf("原始map地址：%p\n", mp)
	change(persons)
	fmt.Println("map被修改，新值为:", persons)
}

func change(mmp map[string]int) {
	fmt.Printf("接收到map的内存地址是：%p\n", &mmp)
	mmp["侯名"] = 666
}
