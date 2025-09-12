package task1

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Age)
}

type Action struct {
	Human
}

func (a Action) Do() {
	a.SayHi()
}
