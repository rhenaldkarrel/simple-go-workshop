package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Dockerized Golang!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on port 8080...")
    http.ListenAndServe(":8080", nil)
}

