package main

import (
	"fmt"
	"time"
)

func main() {
	bufCh := make(chan int, 20)

	go func() {
		for {
			var items []int
			items = append(items, <-bufCh)

			// As we want to batch get maximum 10 items, 9 items remain to get.
			remains := 9

		Remaining:
			for i := 0; i < remains; i++ {
				select {
				case item := <-bufCh:
					items = append(items, item)
				default:
					break Remaining
				}
			}

			// The batch processing. Here we just log output.
			fmt.Println("Items:", items)
		}
	}()

	for i := 0; i < 5; i++ {
		bufCh <- i
		fmt.Println("Push:", i)
	}

	time.Sleep(time.Second)
}
