package main

import "fmt"

type OldPrinter interface {
	Print(s string) string
}

type NewPrinter interface {
	PrintModern(s string) string
}

type OldPrinterImpl struct {
}

func (op *OldPrinterImpl) Print(s string) string {
	return "Old printer" + s
}

type PrinterAdapter struct {
	OldPrinter
}

func (p *PrinterAdapter) PrintModern(s string) string {
	// Здесь мы адаптируем старый метод под новый интерфейс
	result := p.OldPrinter.Print(s)
	return result
}

func main() {
	/*
		Реальный пример, например, для подключения к разным бд, может быть такое, что система ожидает Connect(url string),
			а библиотека предоставляет Open(url string), тогда адаптер поможет нам.

		Плюсы:
			1. Можно использовать старый код, с новым интерфейсом
			2. Упрощает интеграцию сторонних библиотек
			3. Способствует принципу открытости/закрытости (OCP)

		Минусы:
			1. Усложняет код
			2. Снижает производительность
	*/

	oldPrinter := &OldPrinterImpl{}

	// Создаём адаптер, который обернёт старый принтер
	adapter := &PrinterAdapter{
		OldPrinter: oldPrinter,
	}

	modernResult := adapter.PrintModern("Привет, мир!")
	fmt.Println(modernResult)
}
