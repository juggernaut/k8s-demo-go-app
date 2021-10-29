package main

import (
    "net/http"
    "fmt"
    "log"
)

func main() {
    http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Welcome to the AUTOMATED deploy of this app!")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
