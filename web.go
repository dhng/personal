package main

import (
"fmt"
"html/template"
"net/http"
)
func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/index", indexHandler)
    fmt.Println("listening...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/index.html")  // Parse template file.
    t.Execute(w, nil)  // merge.
}