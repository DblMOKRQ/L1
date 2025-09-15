package main

import "fmt"

func reverse(r []rune, i int, j int) {
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
}

func reverseWords(r []rune) string {
	n := len(r)
	reverse(r, 0, n-1)
	start := 0
	for i := 0; i <= n; i++ {
		if i == n || r[i] == ' ' {
			reverse(r, start, i-1)
			start = i + 1
		}
	}
	return string(r)
}

func main() {
	// Можно было указать []byte вместо []rune, но тогда не будет поддержки UTF-8
	// Например, "Привет мир" не будет работать корректно с []byte
	s := "snow dog sun"
	r := []rune(s)
	fmt.Println(reverseWords(r))
}
