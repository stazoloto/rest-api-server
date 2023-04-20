package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Muxer
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/snippet", app.showSnippetHandler)
	mux.HandleFunc("/snippet/create", app.createSnippetHandler)

	// File server for static files
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
