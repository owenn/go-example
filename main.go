package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

type Page struct {
    Title string
    Body  []byte
}


func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    filename := "index.html"
    body, err := ioutil.ReadFile(filename)

    if err != nil {
        return
    }

    w.Header().Set("Content-Type", "text/html")

    fmt.Fprintf(w, "%s", body)
}

func main() {
    http.HandleFunc("/index/", viewHandler)
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8080", nil)
}
