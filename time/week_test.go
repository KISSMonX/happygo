package main

import (
	"fmt"
	"testing"
	"time"
)

func TestWeekNumber(t *testing.T) {
	y, w := time.Now().ISOWeek()
	fmt.Println(y, w)
}
