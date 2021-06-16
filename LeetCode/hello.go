package main

import (
	"fmt"
	"time"
)

func main() {

	format := "2006-01-02 15:04:05"
	timeLeft, _ := time.ParseInLocation(format, "2021-06-03 00:00:00", time.Local)
	timeRight, _ := time.ParseInLocation(format, "2021-06-05 23:59:59", time.Local)

	curTime := time.Now()

	leftDiff := curTime.Sub(timeLeft)
	rightDiff := curTime.Sub(timeRight)
	if leftDiff > 0 && rightDiff < 0 {
		fmt.Println("!!!!")
	}

}
