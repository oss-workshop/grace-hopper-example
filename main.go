package main

import (
	"fmt"
	"log"
	"net/http"
)

const name = "Grace Hopper"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Open Source ♡ %s!", name)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
