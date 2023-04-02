package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet", showSnippetHandler)
	mux.HandleFunc("/snippet/create", createSnippetHandler)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Printf("Starting server at port 3000\n")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
