package main

import (
    "net/http"
    "text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("view.html") //setp 1
    t.Execute(w, "Hello World!") //step 2
}

func main() {

    server := http.Server{
        Addr: "127.0.0.1:8080",
    }
    
    http.HandleFunc("/view", handler)
    server.ListenAndServe()
}