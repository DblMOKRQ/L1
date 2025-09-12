package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func worker(ch <-chan int) {
	//defer wg.Done()
	for num := range ch {
		fmt.Printf("Worker received: %d\n", num)
	}

}

func main() {
	// Примечание: Я думаю можно и, наверное, даже нужно использовать буферизированный канал,
	// чтобы не возникло deadlock, но, вроде, и так работает))
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Please provide a valid positive integer for number of workers")
		return
	}
	//var wg sync.WaitGroup

	ch := make(chan int)

	for i := 0; i < n; i++ {
		//wg.Add(1)
		//go worker(ch, &wg) Можно использовать WaitGroup, перед этим добавить к параметрам функции wg *sync.WaitGroup
		go worker(ch)
	}

	for {
		ch <- rand.Intn(100)
	}

	//close(ch)

}
