package main

import (
	"fmt"
	"sync"
	"time"
)

// Пример с race condition
/*
import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			m[fmt.Sprintf("key%d", i)] = i
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(m)
}
*/
// Пример без race condition
func main() {
	var m sync.Map
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Store(fmt.Sprintf("key%d", i), i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
