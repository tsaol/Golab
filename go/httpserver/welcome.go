package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, I got your request: %s\n", r.URL.Path)
    })

    //static assets
    /*
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    */
    //listen port
    http.ListenAndServe(":80", nil)
}
