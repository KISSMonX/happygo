package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f := bufio.NewReader(os.Stdin)

	for {
		txt, _ := f.ReadString('\n') //定义一行输入的内容分隔符。

		fmt.Println("我: ", txt)
		fmt.Println("AI: ", strings.Replace(txt, "吗?", "!", -1))
	}
}
