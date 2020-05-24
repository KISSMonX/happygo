package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println(GenerateUserID())

}

var (
	initial     = 1
	idLastTime  time.Time
	numlastTime time.Time
	increase    int
	lock        sync.Mutex
)

// getSerialNo 获取序列号
// Get the serial no within a second; lock needed in case multiple orders may get created within a sec
// Alternatively using "atomic", such as atomic.AddInt32
func getSerialNo() int {
	lock.Lock()
	now := time.Now()

	if now.Unix() != idLastTime.Unix() {
		increase = initial
		idLastTime = now
	} else {
		increase++
	}

	lock.Unlock()

	return increase
}

func GenerateUserID() string {
	var idstr string
	t := time.Now().Unix()
	prefix := time.Unix(t, 0).Format("060102")
	idstr += prefix

	// need 0 to pad the left
	idstr += fmt.Sprintf("%03d", getSerialNo())

	return idstr
}
