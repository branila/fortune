package main

import (
	"log"
	"net/http"

	"github.com/branila/fortune/handler"
)

func main() {
	http.HandleFunc("/telegram", handler.Master)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
