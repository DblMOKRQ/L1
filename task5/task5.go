package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func worker(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	// Не стал создавать много горутин, можно посмотреть пример из предыдущего задания
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <time in seconds>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Please provide a valid positive integer")
		return
	}

	ch := make(chan int, 10)
	go worker(ch)
	timeout := time.After(time.Duration(n) * time.Second)
	for {
		select {
		case <-timeout:
			return
		default:
			ch <- rand.Intn(100)
		}
	}

}
