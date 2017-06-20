package main

import (
"fmt"
"net/http"
)

func main() {
    http.HandleFunc("/", index)
    fmt.Println("listening...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}

func index(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "index page")
}
