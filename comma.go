package main

import "fmt"

//вставляет запятые в строковое представление
//неотрицательного десятичного числа

func main() {
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	fmt.Println(s)
	//рекурсивный вызов
	return comma(s[:n-3]) + "," + s[n-3:]
}
