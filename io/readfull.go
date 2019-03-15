package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("0123456789")

	buf := make([]byte, 3)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 10)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		if err != io.ErrUnexpectedEOF {
			fmt.Println("error:", err)
		}
	}
	fmt.Printf("%s\n", longBuf)
}

// output:
// 012
// error: unexpected EOF
// 3456789
