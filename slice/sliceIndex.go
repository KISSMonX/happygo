package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "20190101"
	txt := "20181011..txt"
	fmt.Println(str[0:6])
	fmt.Println(strings.TrimSuffix(txt, ".txt"))
}
