package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "os"
)
//go get -u github.com/gorilla/mux # install dependencies
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wow")
    })

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]
        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    port := ":8000"

    if os.Getenv("OPTIC_API_PORT") != "" {
        port = fmt.Sprintf(":%s", os.Getenv("OPTIC_API_PORT"))
    }

    http.ListenAndServe(port, r)
}
