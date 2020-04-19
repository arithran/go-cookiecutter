package main

import (
	"net/http"
	"{{SERVICE}}/handler"
)

func main() {
	// curl localhost:8080/info
	http.HandleFunc("/info", handler.Info)

	// start {{SERVICE}}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
