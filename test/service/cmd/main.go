package main

import (
    "fmt"
    "net/http"
)

func errorResponse(w http.ResponseWriter, req *http.Request) {

    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/error", errorResponse)

    http.ListenAndServe(":8000", nil)
}
