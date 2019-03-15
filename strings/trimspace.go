package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `{"msg":"bellwether;
	","msg_type":0,"name":"Mike Li","role":10,"sender":"6e925da2527745548df39fd6e1fe51d7"}`

	fmt.Println(strings.Replace(str, "\n", "", -1))
}
