package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func writeWorker(ctx context.Context, wg *sync.WaitGroup, rowChan chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			rowChan <- rand.Intn(100)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func processWorker(ctx context.Context, wg *sync.WaitGroup, rowChan chan int, doneChan chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case x := <-rowChan:
			fmt.Printf("x: %d ", x)
			x = x * 2
			doneChan <- x
		}
	}
}

func readWorker(ctx context.Context, wg *sync.WaitGroup, doneChan chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case x := <-doneChan:
			fmt.Printf("x * 2: %d\n", x)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	rowChan := make(chan int)
	doneChan := make(chan int)
	wg.Add(1)
	go writeWorker(ctx, &wg, rowChan)
	wg.Add(1)
	go processWorker(ctx, &wg, rowChan, doneChan)
	wg.Add(1)
	go readWorker(ctx, &wg, doneChan)

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	close(rowChan)
	close(doneChan)

}
