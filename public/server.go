package main

// Taken from https://github.com/seqsense/pcdeditor/blob/master/examples/serve/main.go

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Println("Serving at http://localhost:8080")
    http.Handle("/", &noCache{Handler: http.FileServer(http.Dir("."))})

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalln(err)
    }
}

type noCache struct {
    http.Handler
}

func (h *noCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "no-cache")
    h.Handler.ServeHTTP(w, r)
}
