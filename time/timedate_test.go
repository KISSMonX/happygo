package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeAdd(t *testing.T) {
	st := time.Now()
	et := st.Add(-10 * time.Minute)
	fmt.Println(st, et)

	fmt.Println("时区: ", time.Now().Location())
	fmt.Println("本地时间: ", time.Now(), "UTC时间: ", time.Now().UTC())
}
