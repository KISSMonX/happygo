package main

import "fmt"

func main() {
	s := `xyz` + `abc`
	ss := "xyz``" + "abc''"
	sss := "xyz"" + `abc'噶十多个' "尬"`
	ssss := '哈' + '🐵'
	fmt.Println(s, ss, sss, ssss)
}
