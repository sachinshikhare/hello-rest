package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
	
    http.HandleFunc("/how-are-you", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "I am okay!")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}