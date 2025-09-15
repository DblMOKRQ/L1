package main

// Выход по условию
/*
func worker(shouldStop *bool) {
	for {
		if *shouldStop {
			fmt.Println("stop")
			return
		}
		fmt.Println("working")
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	shouldStop := false
	go worker(&shouldStop)

	time.Sleep(3 * time.Second)
	shouldStop = true
	time.Sleep(1 * time.Second)
}
*/

// Использование канала для уведомления
/*
func worker(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("stop")
			return
		default:
			fmt.Println("working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan struct{})
	go worker(stopChan)

	time.Sleep(3 * time.Second)
	close(stopChan)
	time.Sleep(1 * time.Second)
}
*/

// Использование контекста для отмены
/*
import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			fmt.Println("working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
*/

// Использование runtime.Goexit()
// runtime.Goexit() немедленно завершает текущую горутину, но в отличие от return,
// он выполняет все отложенные функции перед выходом.
/*
func worker() {
	defer fmt.Println("worker exited")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("stop")
			runtime.Goexit()
		}
		fmt.Println("working", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go worker()

	time.Sleep(5 * time.Second)
}
*/

// Использование panic и recover
/*
func worker() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Goroutine stopped from panic:", r)
		}
	}()
	for i := 0; i < 10; i++ {
		if i == 5 {
			panic("stop")
		}
		fmt.Println("working", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go worker()

	time.Sleep(5 * time.Second)
}
*/

// Таймаут через time.After
/*
func worker(timeout time.Duration) {
	timer := time.After(timeout)
	for {
		select {
		case <-timer:
			fmt.Println("stop due to timeout")
			return
		default:
			fmt.Println("working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	go worker(3 * time.Second)
	time.Sleep(5 * time.Second)
}
*/

// Использование sync.WaitGroup для координации завершения
/*
func worker(stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("worker stopped")
			return
		default:
			fmt.Println("working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	wg.Add(1)
	go worker(stop, &wg)

	time.Sleep(3 * time.Second)
	close(stop)

	wg.Wait()
}
*/

// Выход из main
// Когда функция main() завершается, runtime Go автоматически останавливает все запущенные горутины
/*
func worker() {
	for {
		fmt.Println("working")
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	go worker()
	time.Sleep(3 * time.Second)
}
*/
