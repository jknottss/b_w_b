package main

import "fmt"

//Изменит базовый срез
//Вернет не пустые строки
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // Срез нулевой длины из исходного среза
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
	data2 := []string{"four", "", "six"}
	fmt.Printf("%q\n", nonempty2(data2))
	fmt.Printf("%q\n", data2)
}
