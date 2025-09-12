package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(num)
		}
	}
}

func start(ch chan int, sigChan chan os.Signal) {
	for {
		select {
		case <-sigChan:
			return
		default:
			ch <- rand.Intn(100)
		}
	}
}

func main() {
	// Обоснование: При получение сигнала ctrl+c или любого другого для завершения программы,
	// у нас создается запись в канал, select видит это и вызывает return (28-29 строка),
	// затем мы вызываем cancel для того, чтобы дать горутинам сигнал, что мы закончили,
	// далее ждем завершение всех горутин (52 строка) и закрываем канал.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, ch, &wg)
	}
	start(ch, sigChan)
	cancel()
	wg.Wait()
	close(ch)
	fmt.Println("all goroutines finished")
}
