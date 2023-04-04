package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	addr := flag.String("addr", ":3000", "HTTP address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet", showSnippetHandler)
	mux.HandleFunc("/snippet/create", createSnippetHandler)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Server start at port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	// Open static files
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	//
	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}
	return f, nil
}
