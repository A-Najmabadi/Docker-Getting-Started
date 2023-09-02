package main

import (
    "fmt"
    "net/http"
)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    }

    http.HandleFunc("/", handler)

    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        fmt.Println("error: ", err)
    }
}
