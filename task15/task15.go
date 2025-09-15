package main

import "strings"

var justString string

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

// В данной ситуации мы получаем утечку памяти, так как при justString = v[:100] мы создаем новую строку,
// которая ссылается на первые 100 байт строки v, из-за этого вся память, которая рассчитана на v не будет освобождена,
// пока justString жив.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

// Тут мы создаем копию среза, т.е теперь justString ссылается на копию среза,
// и v может быть спокойно удален, после выхода из функции
func someFuncCorrect() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
}
