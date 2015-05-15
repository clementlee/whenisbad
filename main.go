package main

import (
    "fmt"
    "net/http"
    "html/template"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func handler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[1:]
    err := templates.ExecuteTemplate(w, "index", title)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/", handler)

    //handle static files, perhaps better left to nginx in future?
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)
}
