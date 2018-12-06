package main

import (
	"fmt"
	"strings"
)

func main() {
	sli := make([]string, 0)

	sli = append(sli, "\"2\"")
	sli = append(sli, "\"2\"")

	// fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(sli), "[]"), " ", ",", -1))
	fmt.Printf("%v", strings.Join(sli, ","))
}
