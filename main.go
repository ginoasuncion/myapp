package main

import (
    "encoding/json"
    "net/http"
)

type Message struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(Message{Message: "Hello from myapp!"})
    })
    http.ListenAndServe(":80", nil)
}

