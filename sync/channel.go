package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	go func() {

		ch = make(chan int, 1)
		ch <- 1
		fmt.Printf("ch 写入: %+v\n", ch)
	}()

	go func(ch chan int) {
		fmt.Printf("ch 读取: %+v\n", ch)

		time.Sleep(time.Second)
		fmt.Printf("ch 读取: %+v\n", ch)

		<-ch
		fmt.Printf("ch 读取: %+v\n", ch)
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
