package main

import (
    "fmt"
    "net/http"
    "testing"
)

type Simple struct {
    Greeting string
    Subject  string
    Host     string
}

func SimpleFactory(host string) Simple {
    return Simple{"Hello", "World", host}
}

func handler(w http.ResponseWriter, r *http.Request) {
    simple := SimpleFactory(r.Host)
    fmt.Fprintf(w, "%v", simple)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
