package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Simple struct {
    Greeting string `json:"greeting"`
    Subject  string `json:"subject"`
    Host     string `json:"host"`
}

func SimpleFactory(host string) Simple {
    return Simple{"Hello", "World", host}
}

func handler(w http.ResponseWriter, r *http.Request) {
    simple := SimpleFactory(r.Host)
    w.Header().Set("Content-Type", "application/json")

    jsonOutput, err := json.Marshal(simple)
    if err != nil {
        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(w, string(jsonOutput))
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at http://localhost:4444")
    log.Fatal(http.ListenAndServe(":4444", nil))
}

