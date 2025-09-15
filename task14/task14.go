package main

import "fmt"

func determineType(obj interface{}) {
	switch obj.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	obj := make(chan int)
	determineType(obj)
}
