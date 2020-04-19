package main

import (
	"net/http"
	"{{SERVICE}}/handler"
)

func main() {
	http.HandleFunc("/info", handler.Info)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
