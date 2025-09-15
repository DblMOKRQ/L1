package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	if duration <= 0 {
		return
	}
	t := time.NewTimer(duration)
	<-t.C
}
func main() {
	fmt.Println("start: " + time.Now().Format("15:04:05"))
	sleep(5 * time.Second)
	fmt.Println("end: " + time.Now().Format("15:04:05"))
}
