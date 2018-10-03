package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	p, err := filepath.Abs(*dir)
	if err != nil {
		panic(err)
	}
	log.Printf("Doc root: %s, listening on %q...", p, *listen)
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(p+"/."))))
}
