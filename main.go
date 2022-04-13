package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello services")

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!\n"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
