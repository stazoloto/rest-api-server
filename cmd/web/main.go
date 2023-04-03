package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":3000", "HTTP address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet", showSnippetHandler)
	mux.HandleFunc("/snippet/create", createSnippetHandler)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Server start at port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
