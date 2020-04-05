package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Printf("Serving . on HTTP port: 8000\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}