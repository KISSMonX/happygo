package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	var count int
	go func() {
		ch <- 1
		fmt.Println("写入 ch 成功, 值: ", <-ch)
	}()

	go func() {
		time.Sleep(time.Second)
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}
