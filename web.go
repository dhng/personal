package main

import (
"fmt"
"net/http"
)

func main() {
    http.HandleFunc("/", indexHandler)
    fmt.Println("listening...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}

func index(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "index page")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}