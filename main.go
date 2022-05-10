package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed docs
//go:embed docs/_app/assets
//go:embed docs/_app/assets/pages/*
//go:embed docs/_app/chunks
//go:embed docs/_app/pages
var embeddedFiles embed.FS

func main() {

	fsys, err := fs.Sub(embeddedFiles, "docs")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(fsys)))
	http.ListenAndServe(":8050", mux)
}
