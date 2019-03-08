package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 20)

	ids := make([]int, 0, 10)
	wait := time.After(time.Second)
	go func() {

	Exit:
		for {
			select {
			case id := <-c:
				// fmt.Println("接收:", id)
				ids = append(ids, id)
				if len(ids) == 10 {
					printSlice(ids)
					ids = []int{}
					wait = time.After(time.Second)
				}
			case <-wait:
				fmt.Println("超时")
				break Exit
				wait = time.After(time.Second)
			}
		}
		fmt.Println("跳出协程")
		return
	}()

	for {
		for i := 0; i < 100; i++ {
			// fmt.Println("写入:", i)
			c <- i
			// time.Sleep(time.Second)
		}

		time.Sleep(3 * time.Second)
		break
	}

	fmt.Println("结束")
	// time.Sleep(time.Minute)
}

func printSlice(ids []int) {
	fmt.Println(ids)
}
