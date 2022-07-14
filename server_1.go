package main

import (
	"fmt"
	"log"
	"net/http"
)

// минимальный echo-сервер

func main() {
	http.HandleFunc("/", handler) //каждый запрос будет вызывать функцию "handler"
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
