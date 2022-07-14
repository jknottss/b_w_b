package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Выводит ответ на запрос по заданному URL
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get("http://" + url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body) // Сразу копируем бади в стандартный вывод без перезаписи в переменную
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("resp status : ", resp.Status) // а вот так можно вывести статус ответа
	}

}
