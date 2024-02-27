package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)

    log.Print("Listening on :8080...")
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        log.Fatal(err)
    }
}
